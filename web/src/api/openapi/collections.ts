/**
 * Generated by orval v6.15.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import useSwr from "swr";
import type { Key, SWRConfiguration } from "swr";

import { fetcher } from "../client";

import type {
  CollectionAddPostOKResponse,
  CollectionCreateBody,
  CollectionCreateOKResponse,
  CollectionGetOKResponse,
  CollectionListOKResponse,
  CollectionRemovePostOKResponse,
  CollectionUpdateBody,
  CollectionUpdateOKResponse,
  InternalServerErrorResponse,
  NotFoundResponse,
  UnauthorisedResponse,
} from "./schemas";

type AwaitedInput<T> = PromiseLike<T> | T;

type Awaited<O> = O extends AwaitedInput<infer T> ? T : never;

/**
 * Create a collection for curating posts under the authenticated account.

 */
export const collectionCreate = (
  collectionCreateBody: CollectionCreateBody,
) => {
  return fetcher<CollectionCreateOKResponse>({
    url: `/v1/collections`,
    method: "post",
    headers: { "Content-Type": "application/json" },
    data: collectionCreateBody,
  });
};

/**
 * List all collections using the filtering options.
 */
export const collectionList = () => {
  return fetcher<CollectionListOKResponse>({
    url: `/v1/collections`,
    method: "get",
  });
};

export const getCollectionListKey = () => [`/v1/collections`] as const;

export type CollectionListQueryResult = NonNullable<
  Awaited<ReturnType<typeof collectionList>>
>;
export type CollectionListQueryError =
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useCollectionList = <
  TError = NotFoundResponse | InternalServerErrorResponse,
>(options?: {
  swr?: SWRConfiguration<Awaited<ReturnType<typeof collectionList>>, TError> & {
    swrKey?: Key;
    enabled?: boolean;
  };
}) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false;
  const swrKey =
    swrOptions?.swrKey ?? (() => (isEnabled ? getCollectionListKey() : null));
  const swrFn = () => collectionList();

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
 * Get a collection by its ID. Collections can be public or private so the
response will depend on which account is making the request and if the
target collection is public, private, owned or not owned by the account.

 */
export const collectionGet = (collectionId: string) => {
  return fetcher<CollectionGetOKResponse>({
    url: `/v1/collections/${collectionId}`,
    method: "get",
  });
};

export const getCollectionGetKey = (collectionId: string) =>
  [`/v1/collections/${collectionId}`] as const;

export type CollectionGetQueryResult = NonNullable<
  Awaited<ReturnType<typeof collectionGet>>
>;
export type CollectionGetQueryError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useCollectionGet = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  collectionId: string,
  options?: {
    swr?: SWRConfiguration<
      Awaited<ReturnType<typeof collectionGet>>,
      TError
    > & { swrKey?: Key; enabled?: boolean };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false && !!collectionId;
  const swrKey =
    swrOptions?.swrKey ??
    (() => (isEnabled ? getCollectionGetKey(collectionId) : null));
  const swrFn = () => collectionGet(collectionId);

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
 * Update a collection owned by the authenticated account.
 */
export const collectionUpdate = (
  collectionId: string,
  collectionUpdateBody: CollectionUpdateBody,
) => {
  return fetcher<CollectionUpdateOKResponse>({
    url: `/v1/collections/${collectionId}`,
    method: "patch",
    headers: { "Content-Type": "application/json" },
    data: collectionUpdateBody,
  });
};

/**
 * Add a post to a collection. The collection must be owned by the account
making the request. The post can be any published post of any kind.

 */
export const collectionAddPost = (collectionId: string, postId: string) => {
  return fetcher<CollectionAddPostOKResponse>({
    url: `/v1/collections/${collectionId}/items/${postId}`,
    method: "put",
  });
};

/**
 * Remove a post from a collection. The collection must be owned by the
account making the request.

 */
export const collectionRemovePost = (collectionId: string, postId: string) => {
  return fetcher<CollectionRemovePostOKResponse>({
    url: `/v1/collections/${collectionId}/items/${postId}`,
    method: "delete",
  });
};
