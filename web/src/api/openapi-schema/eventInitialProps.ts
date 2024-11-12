/**
 * Generated by orval v7.2.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
The Storyden API does not adhere to semantic versioning but instead applies a rolling strategy with deprecations and minimal breaking changes. This has been done mainly for a simpler development process and it may be changed to a more fixed versioning strategy in the future. Ultimately, the primary way Storyden tracks versions is dates, there are no set release tags currently.

 * OpenAPI spec version: rolling
 */
import type { EventCapacity } from "./eventCapacity";
import type { EventDescription } from "./eventDescription";
import type { EventLocation } from "./eventLocation";
import type { EventName } from "./eventName";
import type { EventParticipationPolicy } from "./eventParticipationPolicy";
import type { EventSlug } from "./eventSlug";
import type { EventTimeRange } from "./eventTimeRange";
import type { Identifier } from "./identifier";
import type { Metadata } from "./metadata";
import type { PostContent } from "./postContent";
import type { Visibility } from "./visibility";

export interface EventInitialProps {
  capacity?: EventCapacity;
  content: PostContent;
  description?: EventDescription;
  location?: EventLocation;
  meta?: Metadata;
  name: EventName;
  participation_policy: EventParticipationPolicy;
  primary_image_asset_id?: Identifier;
  slug?: EventSlug;
  thread_category_id: Identifier;
  time_range: EventTimeRange;
  visibility: Visibility;
}
