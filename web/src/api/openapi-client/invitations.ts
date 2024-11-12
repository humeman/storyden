/**
 * Generated by orval v7.2.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
The Storyden API does not adhere to semantic versioning but instead applies a rolling strategy with deprecations and minimal breaking changes. This has been done mainly for a simpler development process and it may be changed to a more fixed versioning strategy in the future. Ultimately, the primary way Storyden tracks versions is dates, there are no set release tags currently.

 * OpenAPI spec version: rolling
 */
import useSwr from "swr";
import type { Arguments, Key, SWRConfiguration } from "swr";
import useSWRMutation from "swr/mutation";
import type { SWRMutationConfiguration } from "swr/mutation";

import { fetcher } from "../client";
import type {
  InternalServerErrorResponse,
  InvitationCreateBody,
  InvitationCreateOKResponse,
  InvitationGetOKResponse,
  InvitationListOKResponse,
  InvitationListParams,
  NotFoundResponse,
  UnauthorisedResponse,
} from "../openapi-schema";

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
export const invitationList = (params?: InvitationListParams) => {
  return fetcher<InvitationListOKResponse>({
    url: `/invitations`,
    method: "GET",
    params,
  });
};

export const getInvitationListKey = (params?: InvitationListParams) =>
  [`/invitations`, ...(params ? [params] : [])] as const;

export type InvitationListQueryResult = NonNullable<
  Awaited<ReturnType<typeof invitationList>>
>;
export type InvitationListQueryError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useInvitationList = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  params?: InvitationListParams,
  options?: {
    swr?: SWRConfiguration<
      Awaited<ReturnType<typeof invitationList>>,
      TError
    > & { swrKey?: Key; enabled?: boolean };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false;
  const swrKey =
    swrOptions?.swrKey ??
    (() => (isEnabled ? getInvitationListKey(params) : null));
  const swrFn = () => invitationList(params);

  const query = useSwr<Awaited<ReturnType<typeof swrFn>>, TError>(
    swrKey,
    swrFn,
    swrOptions,
  );

  return {
    swrKey,
    ...query,
  };
};
/**
 * Create an invitation for the authenticated account. Responds with the
invitation data which can be used to construct a public vendor-specific
registration URL using the invitation's identifier which can be used in
calls to registration operations to indicate the account was invited.

 */
export const invitationCreate = (
  invitationCreateBody: InvitationCreateBody,
) => {
  return fetcher<InvitationCreateOKResponse>({
    url: `/invitations`,
    method: "POST",
    headers: { "Content-Type": "application/json" },
    data: invitationCreateBody,
  });
};

export const getInvitationCreateMutationFetcher = () => {
  return (
    _: Key,
    { arg }: { arg: InvitationCreateBody },
  ): Promise<InvitationCreateOKResponse> => {
    return invitationCreate(arg);
  };
};
export const getInvitationCreateMutationKey = () => [`/invitations`] as const;

export type InvitationCreateMutationResult = NonNullable<
  Awaited<ReturnType<typeof invitationCreate>>
>;
export type InvitationCreateMutationError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useInvitationCreate = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(options?: {
  swr?: SWRMutationConfiguration<
    Awaited<ReturnType<typeof invitationCreate>>,
    TError,
    Key,
    InvitationCreateBody,
    Awaited<ReturnType<typeof invitationCreate>>
  > & { swrKey?: string };
}) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey = swrOptions?.swrKey ?? getInvitationCreateMutationKey();
  const swrFn = getInvitationCreateMutationFetcher();

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
/**
 * Retrieve the details of an invitation by its identifier. This endpoint
is publicly accessible and can be used to show invitation details before
the client's registration flow.

 */
export const invitationGet = (invitationId: string) => {
  return fetcher<InvitationGetOKResponse>({
    url: `/invitations/${invitationId}`,
    method: "GET",
  });
};

export const getInvitationGetKey = (invitationId: string) =>
  [`/invitations/${invitationId}`] as const;

export type InvitationGetQueryResult = NonNullable<
  Awaited<ReturnType<typeof invitationGet>>
>;
export type InvitationGetQueryError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useInvitationGet = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  invitationId: string,
  options?: {
    swr?: SWRConfiguration<
      Awaited<ReturnType<typeof invitationGet>>,
      TError
    > & { swrKey?: Key; enabled?: boolean };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false && !!invitationId;
  const swrKey =
    swrOptions?.swrKey ??
    (() => (isEnabled ? getInvitationGetKey(invitationId) : null));
  const swrFn = () => invitationGet(invitationId);

  const query = useSwr<Awaited<ReturnType<typeof swrFn>>, TError>(
    swrKey,
    swrFn,
    swrOptions,
  );

  return {
    swrKey,
    ...query,
  };
};
/**
 * Delete an invitation. After deletion, it cannot be used.
 */
export const invitationDelete = (invitationId: string) => {
  return fetcher<void>({
    url: `/invitations/${invitationId}`,
    method: "DELETE",
  });
};

export const getInvitationDeleteMutationFetcher = (invitationId: string) => {
  return (_: Key, __: { arg: Arguments }): Promise<void> => {
    return invitationDelete(invitationId);
  };
};
export const getInvitationDeleteMutationKey = (invitationId: string) =>
  [`/invitations/${invitationId}`] as const;

export type InvitationDeleteMutationResult = NonNullable<
  Awaited<ReturnType<typeof invitationDelete>>
>;
export type InvitationDeleteMutationError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useInvitationDelete = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  invitationId: string,
  options?: {
    swr?: SWRMutationConfiguration<
      Awaited<ReturnType<typeof invitationDelete>>,
      TError,
      Key,
      Arguments,
      Awaited<ReturnType<typeof invitationDelete>>
    > & { swrKey?: string };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey =
    swrOptions?.swrKey ?? getInvitationDeleteMutationKey(invitationId);
  const swrFn = getInvitationDeleteMutationFetcher(invitationId);

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
