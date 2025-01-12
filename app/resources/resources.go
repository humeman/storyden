package resources

import (
	"go.uber.org/fx"

	"github.com/Southclaws/storyden/app/resources/account/account_querier"
	"github.com/Southclaws/storyden/app/resources/account/account_writer"
	"github.com/Southclaws/storyden/app/resources/account/authentication"
	"github.com/Southclaws/storyden/app/resources/account/email"
	"github.com/Southclaws/storyden/app/resources/account/invitation/invitation_querier"
	"github.com/Southclaws/storyden/app/resources/account/invitation/invitation_writer"
	"github.com/Southclaws/storyden/app/resources/account/notification/notify_querier"
	"github.com/Southclaws/storyden/app/resources/account/notification/notify_writer"
	"github.com/Southclaws/storyden/app/resources/account/role/role_assign"
	"github.com/Southclaws/storyden/app/resources/account/role/role_badge"
	"github.com/Southclaws/storyden/app/resources/account/role/role_querier"
	"github.com/Southclaws/storyden/app/resources/account/role/role_writer"
	"github.com/Southclaws/storyden/app/resources/asset/asset_querier"
	"github.com/Southclaws/storyden/app/resources/asset/asset_writer"
	collection_items "github.com/Southclaws/storyden/app/resources/collection/collection_item"
	"github.com/Southclaws/storyden/app/resources/collection/collection_querier"
	"github.com/Southclaws/storyden/app/resources/collection/collection_writer"
	"github.com/Southclaws/storyden/app/resources/datagraph/hydrate"
	"github.com/Southclaws/storyden/app/resources/event/event_querier"
	"github.com/Southclaws/storyden/app/resources/event/event_writer"
	"github.com/Southclaws/storyden/app/resources/event/participation/participant_querier"
	"github.com/Southclaws/storyden/app/resources/event/participation/participant_writer"
	"github.com/Southclaws/storyden/app/resources/library/node_children"
	"github.com/Southclaws/storyden/app/resources/library/node_querier"
	"github.com/Southclaws/storyden/app/resources/library/node_search"
	"github.com/Southclaws/storyden/app/resources/library/node_traversal"
	"github.com/Southclaws/storyden/app/resources/library/node_writer"
	"github.com/Southclaws/storyden/app/resources/like/like_querier"
	"github.com/Southclaws/storyden/app/resources/like/like_writer"
	"github.com/Southclaws/storyden/app/resources/link/link_querier"
	"github.com/Southclaws/storyden/app/resources/link/link_writer"
	"github.com/Southclaws/storyden/app/resources/post/category"
	"github.com/Southclaws/storyden/app/resources/post/post_search"
	"github.com/Southclaws/storyden/app/resources/post/post_writer"
	"github.com/Southclaws/storyden/app/resources/post/reaction"
	"github.com/Southclaws/storyden/app/resources/post/reply"
	"github.com/Southclaws/storyden/app/resources/post/thread"
	"github.com/Southclaws/storyden/app/resources/post/thread_cache"
	"github.com/Southclaws/storyden/app/resources/profile/follow_querier"
	"github.com/Southclaws/storyden/app/resources/profile/follow_writer"
	"github.com/Southclaws/storyden/app/resources/profile/profile_cache"
	"github.com/Southclaws/storyden/app/resources/profile/profile_search"
	"github.com/Southclaws/storyden/app/resources/question"
	"github.com/Southclaws/storyden/app/resources/settings"
	"github.com/Southclaws/storyden/app/resources/tag/tag_querier"
	"github.com/Southclaws/storyden/app/resources/tag/tag_writer"
)

func Build() fx.Option {
	return fx.Options(
		fx.Provide(
			settings.New,
			account_querier.New,
			account_writer.New,
			email.New,
			role_assign.New,
			role_querier.New,
			role_writer.New,
			role_badge.New,
			invitation_querier.New,
			invitation_writer.New,
			asset_querier.New,
			asset_writer.New,
			authentication.New,
			category.New,
			notify_querier.New,
			notify_writer.New,
			reply.New,
			tag_querier.New,
			tag_writer.New,
			thread.New,
			thread_cache.New,
			reaction.New,
			like_querier.New,
			like_writer.New,
			post_search.New,
			post_writer.New,
			collection_querier.New,
			collection_writer.New,
			collection_items.New,
			node_querier.New,
			node_writer.New,
			node_traversal.New,
			node_children.New,
			node_search.New,
			link_querier.New,
			link_writer.New,
			profile_search.New,
			profile_cache.New,
			follow_writer.New,
			follow_querier.New,
			event_querier.New,
			event_writer.New,
			participant_querier.New,
			participant_writer.New,
			hydrate.New,
			question.New,
		),
	)
}
