/**
 * Generated by orval v6.28.2 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import type { CollectionList } from "./collectionList";
import type { NodeList } from "./nodeList";
import type { PostList } from "./postList";
import type { ThreadList } from "./threadList";

export type LinkWithRefsAllOf = {
  collections: CollectionList;
  nodes: NodeList;
  posts: PostList;
  threads: ThreadList;
};
