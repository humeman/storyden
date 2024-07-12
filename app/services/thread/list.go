package thread

import (
	"context"
	"time"

	"github.com/Southclaws/fault"
	"github.com/Southclaws/fault/fctx"
	"github.com/Southclaws/fault/fmsg"
	"github.com/Southclaws/opt"
	"github.com/rs/xid"

	"github.com/Southclaws/storyden/app/resources/account"
	"github.com/Southclaws/storyden/app/resources/post/thread"
	"github.com/Southclaws/storyden/app/resources/visibility"
)

type Params struct {
	Query         opt.Optional[string]
	CreatedBefore opt.Optional[time.Time]
	UpdatedBefore opt.Optional[time.Time]
	AccountID     opt.Optional[account.AccountID]
	Tags          opt.Optional[[]xid.ID]
	Categories    opt.Optional[[]string]
}

func (s *service) List(ctx context.Context,
	page int,
	size int,
	opts Params,
) (*thread.Result, error) {
	q := []thread.Query{
		// User's drafts are always private so we always filter published only.
		thread.HasStatus(visibility.VisibilityPublished),
		thread.HasNotBeenDeleted(),
	}

	opts.Query.Call(func(value string) { q = append(q, thread.HasKeyword(value)) })
	opts.CreatedBefore.Call(func(value time.Time) { q = append(q, thread.HasCreatedDateBefore(value)) })
	opts.UpdatedBefore.Call(func(value time.Time) { q = append(q, thread.HasUpdatedDateBefore(value)) })
	opts.AccountID.Call(func(a account.AccountID) { q = append(q, thread.HasAuthor(a)) })
	opts.Tags.Call(func(a []xid.ID) { q = append(q, thread.HasTags(a)) })
	opts.Categories.Call(func(a []string) { q = append(q, thread.HasCategories(a)) })

	thr, err := s.thread_repo.List(ctx, page, size, q...)
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx), fmsg.With("failed to list threads"))
	}

	return thr, nil
}
