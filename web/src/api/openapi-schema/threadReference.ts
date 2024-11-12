/**
 * Generated by orval v7.2.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
The Storyden API does not adhere to semantic versioning but instead applies a rolling strategy with deprecations and minimal breaking changes. This has been done mainly for a simpler development process and it may be changed to a more fixed versioning strategy in the future. Ultimately, the primary way Storyden tracks versions is dates, there are no set release tags currently.

 * OpenAPI spec version: rolling
 */
import type { CommonProperties } from "./commonProperties";
import type { PostProps } from "./postProps";
import type { PostReferenceProps } from "./postReferenceProps";
import type { ThreadReferenceProps } from "./threadReferenceProps";

/**
 * A thread reference includes most of the information about a thread but
does not include the posts within the thread. Useful for rendering large
lists of threads or other situations when you don't need the full data.

 */
export type ThreadReference = CommonProperties &
  PostReferenceProps &
  PostProps &
  ThreadReferenceProps;
