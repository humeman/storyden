/**
 * Generated by orval v6.30.2 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import type { CommonProperties } from "./commonProperties";
import type { PostProps } from "./postProps";

/**
 * A post represents a temporal piece of content, it can be a thread, or a
reply to a thread or something else such as a blog, announcement, etc.
Post is used in generic use-cases where it may not matter whether you
want a thread or a reply, such as search results or recommendations.

 */
export type Post = CommonProperties & PostProps;
