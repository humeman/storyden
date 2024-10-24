/**
 * Generated by orval v6.31.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
The Storyden API does not adhere to semantic versioning but instead applies a rolling strategy with deprecations and minimal breaking changes. This has been done mainly for a simpler development process and it may be changed to a more fixed versioning strategy in the future. Ultimately, the primary way Storyden tracks versions is dates, there are no set release tags currently.

 * OpenAPI spec version: rolling
 */
import type { CategoryReference } from "./categoryReference";
import type { LinkReference } from "./linkReference";
import type { ReplyStatus } from "./replyStatus";
import type { TagList } from "./tagList";

export interface ThreadReferenceProps {
  category: CategoryReference;
  link?: LinkReference;
  /** Whether the thread is pinned in this category. */
  pinned: boolean;
  reply_status: ReplyStatus;
  tags: TagList;
}
