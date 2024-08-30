/**
 * Generated by orval v6.31.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
The Storyden API does not adhere to semantic versioning but instead applies a rolling strategy with deprecations and minimal breaking changes. This has been done mainly for a simpler development process and it may be changed to a more fixed versioning strategy in the future. Ultimately, the primary way Storyden tracks versions is dates, there are no set release tags currently.

 * OpenAPI spec version: rolling
 */
import type { AccountBio } from "./accountBio";
import type { AccountEmailAddressList } from "./accountEmailAddressList";
import type { AccountHandle } from "./accountHandle";
import type { AccountName } from "./accountName";
import type { AccountVerifiedStatus } from "./accountVerifiedStatus";
import type { Metadata } from "./metadata";
import type { ProfileExternalLinkList } from "./profileExternalLinkList";
import type { TagList } from "./tagList";

export interface AccountCommonProps {
  admin: boolean;
  bio: AccountBio;
  email_addresses: AccountEmailAddressList;
  handle: AccountHandle;
  interests?: TagList;
  links: ProfileExternalLinkList;
  meta: Metadata;
  name: AccountName;
  verified_status: AccountVerifiedStatus;
}
