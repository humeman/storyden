/**
 * Generated by orval v6.15.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import type { CollectionCommonProps } from "./collectionCommonProps";
import type { CommonProperties } from "./commonProperties";

/**
 * A collection is a group of threads owned by a user. It allows users to
curate their own lists of content from the site. Collections can only
contain root level posts (threads) with titles and slugs to link to.

 */
export type Collection = CommonProperties & CollectionCommonProps;
