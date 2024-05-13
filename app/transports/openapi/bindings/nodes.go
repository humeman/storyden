package bindings

import (
	"context"

	"github.com/Southclaws/dt"
	"github.com/Southclaws/fault"
	"github.com/Southclaws/fault/fctx"
	"github.com/Southclaws/fault/ftag"
	"github.com/Southclaws/opt"
	"github.com/rs/xid"
	"github.com/samber/lo"

	"github.com/Southclaws/storyden/app/resources/account"
	"github.com/Southclaws/storyden/app/resources/asset"
	"github.com/Southclaws/storyden/app/resources/datagraph"
	"github.com/Southclaws/storyden/app/resources/datagraph/node_traversal"
	"github.com/Southclaws/storyden/app/resources/post"
	"github.com/Southclaws/storyden/app/services/authentication/session"
	node_svc "github.com/Southclaws/storyden/app/services/node"
	"github.com/Southclaws/storyden/app/services/node/node_visibility"
	"github.com/Southclaws/storyden/app/services/nodetree"
	"github.com/Southclaws/storyden/app/transports/openapi"
)

type Nodes struct {
	ar    account.Repository
	cs    node_svc.Manager
	nv    *node_visibility.Controller
	ntree nodetree.Graph
	ntr   node_traversal.Repository
}

func NewNodes(
	ar account.Repository,
	cs node_svc.Manager,
	nv *node_visibility.Controller,
	ntree nodetree.Graph,
	ntr node_traversal.Repository,
) Nodes {
	return Nodes{
		ar:    ar,
		cs:    cs,
		nv:    nv,
		ntree: ntree,
		ntr:   ntr,
	}
}

func (c *Nodes) NodeCreate(ctx context.Context, request openapi.NodeCreateRequestObject) (openapi.NodeCreateResponseObject, error) {
	session, err := session.GetAccountID(ctx)
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	vis, err := opt.MapErr(opt.NewPtr(request.Body.Visibility), deserialiseVisibility)
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	node, err := c.cs.Create(ctx,
		session,
		request.Body.Name,
		request.Body.Slug,
		request.Body.Description,
		node_svc.Partial{
			Content:    opt.NewPtr(request.Body.Content),
			Properties: opt.NewPtr(request.Body.Properties),
			URL:        opt.NewPtr(request.Body.Url),
			AssetsAdd:  opt.NewPtrMap(request.Body.AssetIds, deserialiseAssetIDs),
			Parent:     opt.NewPtrMap(request.Body.Parent, deserialiseNodeSlug),
			Visibility: vis,
		},
	)
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	return openapi.NodeCreate200JSONResponse{
		NodeCreateOKJSONResponse: openapi.NodeCreateOKJSONResponse(serialiseNode(node)),
	}, nil
}

func (c *Nodes) NodeList(ctx context.Context, request openapi.NodeListRequestObject) (openapi.NodeListResponseObject, error) {
	acc, err := opt.MapErr(session.GetOptAccountID(ctx), func(aid account.AccountID) (*account.Account, error) {
		return c.ar.GetByID(ctx, aid)
	})
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	var cs []*datagraph.Node

	opts := []node_traversal.Filter{}

	if v := request.Params.Author; v != nil {
		opts = append(opts, node_traversal.WithOwner(*v))
	}

	visibilities, err := opt.MapErr(opt.NewPtr(request.Params.Visibility), deserialiseVisibilityList)
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	if a, ok := acc.Get(); ok {
		// NOTE: We do not want to allow anyone to request ANY node that is
		// not published, but we also want to allow admins to request nodes
		// that are in review. So we need to check if the requesting account is
		// an admin and if they are not, automatically add a WithOwner filter.

		if v, ok := visibilities.Get(); ok {
			opts = append(opts, node_traversal.WithVisibility(v...))

			if lo.Contains(v, post.VisibilityDraft) {
				// If the result is to contain drafts, only show the account's.
				opts = append(opts, node_traversal.WithOwner(a.Handle))
			} else if lo.Contains(v, post.VisibilityReview) {
				// If the result is to contain nodes that are in-review, then
				// we need to check if the requesting account is an admin first.
				if !a.Admin {
					opts = append(opts, node_traversal.WithOwner(a.Handle))
				}
			}
		}
	} else {
		// When the request is not made by an authenticated account, we do not
		// permit any visibility other than "published".

		opts = append(opts, node_traversal.WithVisibility(post.VisibilityPublished))
	}

	if id := request.Params.NodeId; id != nil {
		cid, err := xid.FromString(*id)
		if err != nil {
			return nil, fault.Wrap(err, fctx.With(ctx), ftag.With(ftag.InvalidArgument))
		}

		cs, err = c.ntr.Subtree(ctx, datagraph.NodeID(cid), opts...)
		if err != nil {
			return nil, fault.Wrap(err, fctx.With(ctx))
		}
	} else {
		cs, err = c.ntr.Root(ctx, opts...)
		if err != nil {
			return nil, fault.Wrap(err, fctx.With(ctx))
		}
	}

	return openapi.NodeList200JSONResponse{
		NodeListOKJSONResponse: openapi.NodeListOKJSONResponse{
			Nodes: dt.Map(cs, serialiseNode),
		},
	}, nil
}

func (c *Nodes) NodeGet(ctx context.Context, request openapi.NodeGetRequestObject) (openapi.NodeGetResponseObject, error) {
	node, err := c.cs.Get(ctx, datagraph.NodeSlug(request.NodeSlug))
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	return openapi.NodeGet200JSONResponse{
		NodeGetOKJSONResponse: openapi.NodeGetOKJSONResponse(serialiseNodeWithItems(node)),
	}, nil
}

func (c *Nodes) NodeUpdate(ctx context.Context, request openapi.NodeUpdateRequestObject) (openapi.NodeUpdateResponseObject, error) {
	node, err := c.cs.Update(ctx, datagraph.NodeSlug(request.NodeSlug), node_svc.Partial{
		Name:        opt.NewPtr(request.Body.Name),
		Slug:        opt.NewPtr(request.Body.Slug),
		AssetsAdd:   opt.NewPtrMap(request.Body.AssetIds, deserialiseAssetIDs),
		URL:         opt.NewPtr(request.Body.Url),
		Description: opt.NewPtr(request.Body.Description),
		Content:     opt.NewPtr(request.Body.Content),
		Parent:      opt.NewPtrMap(request.Body.Parent, deserialiseNodeSlug),
		Properties:  opt.NewPtr(request.Body.Properties),
	})
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	return openapi.NodeUpdate200JSONResponse{
		NodeUpdateOKJSONResponse: openapi.NodeUpdateOKJSONResponse(serialiseNode(node)),
	}, nil
}

func (c *Nodes) NodeUpdateVisibility(ctx context.Context, request openapi.NodeUpdateVisibilityRequestObject) (openapi.NodeUpdateVisibilityResponseObject, error) {
	v, err := post.NewVisibility(string(request.Body.Visibility))
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx), ftag.With(ftag.InvalidArgument))
	}

	node, err := c.nv.ChangeVisibility(ctx, datagraph.NodeSlug(request.NodeSlug), v)
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	return openapi.NodeUpdateVisibility200JSONResponse{
		NodeUpdateOKJSONResponse: openapi.NodeUpdateOKJSONResponse(serialiseNode(node)),
	}, nil
}

func (c *Nodes) NodeDelete(ctx context.Context, request openapi.NodeDeleteRequestObject) (openapi.NodeDeleteResponseObject, error) {
	destinationNode, err := c.cs.Delete(ctx, datagraph.NodeSlug(request.NodeSlug), node_svc.DeleteOptions{
		MoveTo: opt.NewPtr((*datagraph.NodeSlug)(request.Params.TargetNode)),
		Nodes:  opt.NewPtr(request.Params.MoveChildNodes).OrZero(),
	})
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	return openapi.NodeDelete200JSONResponse{
		NodeDeleteOKJSONResponse: openapi.NodeDeleteOKJSONResponse{
			Destination: opt.Map(opt.NewPtr(destinationNode), func(in datagraph.Node) openapi.Node {
				return serialiseNode(&in)
			}).Ptr(),
		},
	}, nil
}

func (c *Nodes) NodeAddAsset(ctx context.Context, request openapi.NodeAddAssetRequestObject) (openapi.NodeAddAssetResponseObject, error) {
	id := openapi.ParseID(request.AssetId)

	node, err := c.cs.Update(ctx, datagraph.NodeSlug(request.NodeSlug), node_svc.Partial{
		AssetsAdd: opt.New([]asset.AssetID{id}),
	})
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	return openapi.NodeAddAsset200JSONResponse{
		NodeUpdateOKJSONResponse: openapi.NodeUpdateOKJSONResponse(serialiseNode(node)),
	}, nil
}

func (c *Nodes) NodeRemoveAsset(ctx context.Context, request openapi.NodeRemoveAssetRequestObject) (openapi.NodeRemoveAssetResponseObject, error) {
	id := openapi.ParseID(request.AssetId)

	node, err := c.cs.Update(ctx, datagraph.NodeSlug(request.NodeSlug), node_svc.Partial{
		AssetsRemove: opt.New([]asset.AssetID{id}),
	})
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	return openapi.NodeRemoveAsset200JSONResponse{
		NodeUpdateOKJSONResponse: openapi.NodeUpdateOKJSONResponse(serialiseNode(node)),
	}, nil
}

func (c *Nodes) NodeAddNode(ctx context.Context, request openapi.NodeAddNodeRequestObject) (openapi.NodeAddNodeResponseObject, error) {
	node, err := c.ntree.Move(ctx, datagraph.NodeSlug(request.NodeSlugChild), datagraph.NodeSlug(request.NodeSlug))
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	return openapi.NodeAddNode200JSONResponse{
		NodeAddChildOKJSONResponse: openapi.NodeAddChildOKJSONResponse(serialiseNode(node)),
	}, nil
}

func (c *Nodes) NodeRemoveNode(ctx context.Context, request openapi.NodeRemoveNodeRequestObject) (openapi.NodeRemoveNodeResponseObject, error) {
	node, err := c.ntree.Sever(ctx, datagraph.NodeSlug(request.NodeSlugChild), datagraph.NodeSlug(request.NodeSlug))
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	return openapi.NodeRemoveNode200JSONResponse{
		NodeRemoveChildOKJSONResponse: openapi.NodeRemoveChildOKJSONResponse(serialiseNode(node)),
	}, nil
}

func serialiseNode(in *datagraph.Node) openapi.Node {
	return openapi.Node{
		Id:          in.ID.String(),
		CreatedAt:   in.CreatedAt,
		UpdatedAt:   in.UpdatedAt,
		Name:        in.Name,
		Slug:        in.Slug,
		Assets:      dt.Map(in.Assets, serialiseAssetReference),
		Link:        opt.Map(in.Links.Latest(), serialiseLink).Ptr(),
		Description: in.Description,
		Content:     in.Content.Ptr(),
		Owner:       serialiseProfileReference(in.Owner),
		Parent:      opt.Map(in.Parent, serialiseNode).Ptr(),
		Visibility:  serialiseVisibility(in.Visibility),
		Properties:  in.Properties,
	}
}

func serialiseNodeWithItems(in *datagraph.Node) openapi.NodeWithChildren {
	return openapi.NodeWithChildren{
		Id:          in.ID.String(),
		CreatedAt:   in.CreatedAt,
		UpdatedAt:   in.UpdatedAt,
		Name:        in.Name,
		Slug:        in.Slug,
		Assets:      dt.Map(in.Assets, serialiseAssetReference),
		Link:        opt.Map(in.Links.Latest(), serialiseLink).Ptr(),
		Description: in.Description,
		Content:     in.Content.Ptr(),
		Owner:       serialiseProfileReference(in.Owner),
		Parent:      opt.Map(in.Parent, serialiseNode).Ptr(),
		Visibility:  serialiseVisibility(in.Visibility),
		Properties:  in.Properties,
		Children:    dt.Map(in.Nodes, serialiseNode),
	}
}

func deserialiseNodeSlug(in string) datagraph.NodeSlug {
	return datagraph.NodeSlug(in)
}
