/**
 * Generated by orval v7.2.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
The Storyden API does not adhere to semantic versioning but instead applies a rolling strategy with deprecations and minimal breaking changes. This has been done mainly for a simpler development process and it may be changed to a more fixed versioning strategy in the future. Ultimately, the primary way Storyden tracks versions is dates, there are no set release tags currently.

 * OpenAPI spec version: rolling
 */
import type { AccountHandle } from "./accountHandle";
import type { CategorySlugList } from "./categorySlugList";
import type { PaginationQueryParameter } from "./paginationQueryParameter";
import type { SearchQueryParameter } from "./searchQueryParameter";
import type { TagListIDs } from "./tagListIDs";
import type { VisibilityParamParameter } from "./visibilityParamParameter";

export type ThreadListParams = {
  /**
   * Search query string.
   */
  q?: SearchQueryParameter;
  /**
   * Pagination query parameters.
   */
  page?: PaginationQueryParameter;
  /**
   * Show only results creeated by this user.
   */
  author?: AccountHandle;
  /**
 * Filter content with specific visibility values. Note that by default,
only published items are returned. When 'draft' is specified, only
drafts owned by the requesting account are included. When 'review' is
specified, the request will fail if the requesting account does not have
the necessary permission to view in-review items.

 */
  visibility?: VisibilityParamParameter;
  /**
   * Show only results with these tags
   */
  tags?: TagListIDs;
  /**
   * Show only results with these categories
   */
  categories?: CategorySlugList;
};
