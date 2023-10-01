/**
 * Generated by orval v6.17.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import type { Identifier } from "./identifier";
import type { Metadata } from "./metadata";
import type { PostContent } from "./postContent";
import type { TagListIDs } from "./tagListIDs";
import type { ThreadStatus } from "./threadStatus";
import type { ThreadTitle } from "./threadTitle";
import type { Url } from "./url";

export interface ThreadMutableProps {
  title?: ThreadTitle;
  body?: PostContent;
  tags?: TagListIDs;
  meta?: Metadata;
  category?: Identifier;
  status?: ThreadStatus;
  url?: Url;
}
