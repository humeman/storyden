/**
 * Generated by orval v6.31.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
The Storyden API does not adhere to semantic versioning but instead applies a rolling strategy with deprecations and minimal breaking changes. This has been done mainly for a simpler development process and it may be changed to a more fixed versioning strategy in the future. Ultimately, the primary way Storyden tracks versions is dates, there are no set release tags currently.

 * OpenAPI spec version: rolling
 */
import type { Identifier } from "./identifier";
import type { TagName } from "./tagName";

/**
 * A tag which can link profiles or threads via interests.
 */
export interface Tag {
  colour: string;
  id: Identifier;
  name: TagName;
}
