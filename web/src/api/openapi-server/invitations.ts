/**
 * Generated by orval v7.2.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
The Storyden API does not adhere to semantic versioning but instead applies a rolling strategy with deprecations and minimal breaking changes. This has been done mainly for a simpler development process and it may be changed to a more fixed versioning strategy in the future. Ultimately, the primary way Storyden tracks versions is dates, there are no set release tags currently.

 * OpenAPI spec version: rolling
 */
import type {
  InvitationCreateBody,
  InvitationCreateOKResponse,
  InvitationGetOKResponse,
  InvitationListOKResponse,
  InvitationListParams,
} from "../openapi-schema";
import { fetcher } from "../server";

/**
 * Retrieve all invitations for the authenticated account. This endpoint
is useful for showing the user which invitations they have sent out and
which ones have been accepted.

If the requesting account is not an admin, the account_id query param 
must be equal to the ID of the requesting session account's ID.

If the requesting account is an admin, the account_id query param may
be used to retrieve invitations for a specific account. Otherwise, the
endpoint will return all invitations for all accounts.

 */
export type invitationListResponse = {
  data: InvitationListOKResponse;
  status: number;
};

export const getInvitationListUrl = (params?: InvitationListParams) => {
  const normalizedParams = new URLSearchParams();

  Object.entries(params || {}).forEach(([key, value]) => {
    if (value !== undefined) {
      normalizedParams.append(key, value === null ? "null" : value.toString());
    }
  });

  return normalizedParams.size
    ? `/invitations?${normalizedParams.toString()}`
    : `/invitations`;
};

export const invitationList = async (
  params?: InvitationListParams,
  options?: RequestInit,
): Promise<invitationListResponse> => {
  return fetcher<Promise<invitationListResponse>>(
    getInvitationListUrl(params),
    {
      ...options,
      method: "GET",
    },
  );
};

/**
 * Create an invitation for the authenticated account. Responds with the
invitation data which can be used to construct a public vendor-specific
registration URL using the invitation's identifier which can be used in
calls to registration operations to indicate the account was invited.

 */
export type invitationCreateResponse = {
  data: InvitationCreateOKResponse;
  status: number;
};

export const getInvitationCreateUrl = () => {
  return `/invitations`;
};

export const invitationCreate = async (
  invitationCreateBody: InvitationCreateBody,
  options?: RequestInit,
): Promise<invitationCreateResponse> => {
  return fetcher<Promise<invitationCreateResponse>>(getInvitationCreateUrl(), {
    ...options,
    method: "POST",
    headers: { "Content-Type": "application/json", ...options?.headers },
    body: JSON.stringify(invitationCreateBody),
  });
};

/**
 * Retrieve the details of an invitation by its identifier. This endpoint
is publicly accessible and can be used to show invitation details before
the client's registration flow.

 */
export type invitationGetResponse = {
  data: InvitationGetOKResponse;
  status: number;
};

export const getInvitationGetUrl = (invitationId: string) => {
  return `/invitations/${invitationId}`;
};

export const invitationGet = async (
  invitationId: string,
  options?: RequestInit,
): Promise<invitationGetResponse> => {
  return fetcher<Promise<invitationGetResponse>>(
    getInvitationGetUrl(invitationId),
    {
      ...options,
      method: "GET",
    },
  );
};

/**
 * Delete an invitation. After deletion, it cannot be used.
 */
export type invitationDeleteResponse = {
  data: void;
  status: number;
};

export const getInvitationDeleteUrl = (invitationId: string) => {
  return `/invitations/${invitationId}`;
};

export const invitationDelete = async (
  invitationId: string,
  options?: RequestInit,
): Promise<invitationDeleteResponse> => {
  return fetcher<Promise<invitationDeleteResponse>>(
    getInvitationDeleteUrl(invitationId),
    {
      ...options,
      method: "DELETE",
    },
  );
};
