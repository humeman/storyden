package bindings

import (
	"context"

	"github.com/Southclaws/dt"
	"github.com/Southclaws/fault"
	"github.com/Southclaws/fault/fctx"
	"github.com/Southclaws/opt"

	"github.com/Southclaws/storyden/app/resources/collection"
	"github.com/Southclaws/storyden/app/resources/post"
	"github.com/Southclaws/storyden/app/services/authentication/session"
	collection_svc "github.com/Southclaws/storyden/app/services/collection"
	"github.com/Southclaws/storyden/internal/openapi"
)

type Collections struct {
	collection_repo collection.Repository
	collection_svc  collection_svc.Service
}

func NewCollections(
	collection_repo collection.Repository,
	collection_svc collection_svc.Service,
) Collections {
	return Collections{
		collection_repo: collection_repo,
		collection_svc:  collection_svc,
	}
}

func (i *Collections) CollectionCreate(ctx context.Context, request openapi.CollectionCreateRequestObject) (openapi.CollectionCreateResponseObject, error) {
	accountID, err := session.GetAccountID(ctx)
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	coll, err := i.collection_repo.Create(ctx, accountID, request.Body.Name, request.Body.Description)
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	return openapi.CollectionCreate200JSONResponse{
		CollectionCreateOKJSONResponse: openapi.CollectionCreateOKJSONResponse(serialiseCollection(coll)),
	}, nil
}

func (i *Collections) CollectionList(ctx context.Context, request openapi.CollectionListRequestObject) (openapi.CollectionListResponseObject, error) {
	colls, err := i.collection_repo.List(ctx)
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	list := dt.Map(colls, serialiseCollection)

	return openapi.CollectionList200JSONResponse{
		CollectionListOKJSONResponse: openapi.CollectionListOKJSONResponse{
			Collections: list,
		},
	}, nil
}

func (i *Collections) CollectionGet(ctx context.Context, request openapi.CollectionGetRequestObject) (openapi.CollectionGetResponseObject, error) {
	coll, err := i.collection_repo.Get(ctx, collection.CollectionID(deserialiseID(request.CollectionId)))
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	return openapi.CollectionGet200JSONResponse{
		CollectionGetOKJSONResponse: openapi.CollectionGetOKJSONResponse(serialiseCollectionWithItems(coll)),
	}, nil
}

func (i *Collections) CollectionUpdate(ctx context.Context, request openapi.CollectionUpdateRequestObject) (openapi.CollectionUpdateResponseObject, error) {
	c, err := i.collection_svc.Update(ctx,
		collection.CollectionID(deserialiseID(request.CollectionId)),
		collection_svc.Partial{
			Name:        opt.NewPtr(request.Body.Name),
			Description: opt.NewPtr(request.Body.Description),
		})
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	return openapi.CollectionUpdate200JSONResponse{
		CollectionUpdateOKJSONResponse: openapi.CollectionUpdateOKJSONResponse(serialiseCollection(c)),
	}, nil
}

func (i *Collections) CollectionAddPost(ctx context.Context, request openapi.CollectionAddPostRequestObject) (openapi.CollectionAddPostResponseObject, error) {
	c, err := i.collection_svc.Add(ctx,
		collection.CollectionID(deserialiseID(request.CollectionId)),
		post.ID(deserialiseID(request.PostId)))
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	return openapi.CollectionAddPost200JSONResponse{
		CollectionAddPostOKJSONResponse: openapi.CollectionAddPostOKJSONResponse(serialiseCollection(c)),
	}, nil
}

func (i *Collections) CollectionRemovePost(ctx context.Context, request openapi.CollectionRemovePostRequestObject) (openapi.CollectionRemovePostResponseObject, error) {
	c, err := i.collection_svc.Remove(ctx,
		collection.CollectionID(deserialiseID(request.CollectionId)),
		post.ID(deserialiseID(request.PostId)))
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	return openapi.CollectionRemovePost200JSONResponse{
		CollectionRemovePostOKJSONResponse: openapi.CollectionRemovePostOKJSONResponse(serialiseCollection(c)),
	}, nil
}

func serialiseCollection(in *collection.Collection) openapi.Collection {
	return openapi.Collection{
		Id:          in.ID.String(),
		CreatedAt:   in.CreatedAt,
		UpdatedAt:   in.UpdatedAt,
		Owner:       serialiseProfileReference(in.Owner),
		Name:        in.Name,
		Description: in.Description,
	}
}

func serialiseCollectionWithItems(in *collection.Collection) openapi.CollectionWithItems {
	return openapi.CollectionWithItems{
		Id:          in.ID.String(),
		CreatedAt:   in.CreatedAt,
		UpdatedAt:   in.UpdatedAt,
		Owner:       serialiseProfileReference(in.Owner),
		Name:        in.Name,
		Description: in.Description,
		Items:       dt.Map(in.Items, serialiseCollectionItem),
	}
}

func serialiseCollectionItem(in *collection.Item) openapi.CollectionItem {
	return openapi.CollectionItem{
		Id:        in.ID.String(),
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
		Slug:      in.Slug,
		Author:    serialiseProfileReference(in.Author),
		Title:     in.Title,
		Short:     in.Short,
	}
}
