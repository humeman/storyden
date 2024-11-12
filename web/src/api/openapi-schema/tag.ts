/**
 * Generated by orval v7.2.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
The Storyden API does not adhere to semantic versioning but instead applies a rolling strategy with deprecations and minimal breaking changes. This has been done mainly for a simpler development process and it may be changed to a more fixed versioning strategy in the future. Ultimately, the primary way Storyden tracks versions is dates, there are no set release tags currently.

 * OpenAPI spec version: rolling
 */
import type { TagProps } from "./tagProps";
import type { TagReferenceProps } from "./tagReferenceProps";

/**
 * A tag is a label that can be applied to posts or pages to organise 
related content. They can be used to filter and search for content.
The Tag schema provides all the data for a tag including its items, so
it's quite a heavy object if referencing a lot of items. For a lighter
weight version, use a TagReference for use-cases such as tag searches.

 */
export type Tag = TagReferenceProps & TagProps;
