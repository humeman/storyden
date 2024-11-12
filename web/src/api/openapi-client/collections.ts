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
  CollectionAddNodeOKResponse,
  CollectionAddPostOKResponse,
  CollectionCreateBody,
  CollectionCreateOKResponse,
  CollectionGetOKResponse,
  CollectionListOKResponse,
  CollectionListParams,
  CollectionRemoveNodeOKResponse,
  CollectionRemovePostOKResponse,
  CollectionUpdateBody,
  CollectionUpdateOKResponse,
  InternalServerErrorResponse,
  NotFoundResponse,
  UnauthorisedResponse,
} from "../openapi-schema";

/**
 * Create a collection for curating posts under the authenticated account.

 */
export const collectionCreate = (
  collectionCreateBody: CollectionCreateBody,
) => {
  return fetcher<CollectionCreateOKResponse>({
    url: `/collections`,
    method: "POST",
    headers: { "Content-Type": "application/json" },
    data: collectionCreateBody,
  });
};

export const getCollectionCreateMutationFetcher = () => {
  return (
    _: Key,
    { arg }: { arg: CollectionCreateBody },
  ): Promise<CollectionCreateOKResponse> => {
    return collectionCreate(arg);
  };
};
export const getCollectionCreateMutationKey = () => [`/collections`] as const;

export type CollectionCreateMutationResult = NonNullable<
  Awaited<ReturnType<typeof collectionCreate>>
>;
export type CollectionCreateMutationError =
  | UnauthorisedResponse
  | InternalServerErrorResponse;

export const useCollectionCreate = <
  TError = UnauthorisedResponse | InternalServerErrorResponse,
>(options?: {
  swr?: SWRMutationConfiguration<
    Awaited<ReturnType<typeof collectionCreate>>,
    TError,
    Key,
    CollectionCreateBody,
    Awaited<ReturnType<typeof collectionCreate>>
  > & { swrKey?: string };
}) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey = swrOptions?.swrKey ?? getCollectionCreateMutationKey();
  const swrFn = getCollectionCreateMutationFetcher();

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
/**
 * List all collections using the filtering options.
 */
export const collectionList = (params?: CollectionListParams) => {
  return fetcher<CollectionListOKResponse>({
    url: `/collections`,
    method: "GET",
    params,
  });
};

export const getCollectionListKey = (params?: CollectionListParams) =>
  [`/collections`, ...(params ? [params] : [])] as const;

export type CollectionListQueryResult = NonNullable<
  Awaited<ReturnType<typeof collectionList>>
>;
export type CollectionListQueryError =
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useCollectionList = <
  TError = NotFoundResponse | InternalServerErrorResponse,
>(
  params?: CollectionListParams,
  options?: {
    swr?: SWRConfiguration<
      Awaited<ReturnType<typeof collectionList>>,
      TError
    > & { swrKey?: Key; enabled?: boolean };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false;
  const swrKey =
    swrOptions?.swrKey ??
    (() => (isEnabled ? getCollectionListKey(params) : null));
  const swrFn = () => collectionList(params);

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
export const collectionGet = (collectionMark: string) => {
  return fetcher<CollectionGetOKResponse>({
    url: `/collections/${collectionMark}`,
    method: "GET",
  });
};

export const getCollectionGetKey = (collectionMark: string) =>
  [`/collections/${collectionMark}`] as const;

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
  collectionMark: string,
  options?: {
    swr?: SWRConfiguration<
      Awaited<ReturnType<typeof collectionGet>>,
      TError
    > & { swrKey?: Key; enabled?: boolean };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false && !!collectionMark;
  const swrKey =
    swrOptions?.swrKey ??
    (() => (isEnabled ? getCollectionGetKey(collectionMark) : null));
  const swrFn = () => collectionGet(collectionMark);

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
  collectionMark: string,
  collectionUpdateBody: CollectionUpdateBody,
) => {
  return fetcher<CollectionUpdateOKResponse>({
    url: `/collections/${collectionMark}`,
    method: "PATCH",
    headers: { "Content-Type": "application/json" },
    data: collectionUpdateBody,
  });
};

export const getCollectionUpdateMutationFetcher = (collectionMark: string) => {
  return (
    _: Key,
    { arg }: { arg: CollectionUpdateBody },
  ): Promise<CollectionUpdateOKResponse> => {
    return collectionUpdate(collectionMark, arg);
  };
};
export const getCollectionUpdateMutationKey = (collectionMark: string) =>
  [`/collections/${collectionMark}`] as const;

export type CollectionUpdateMutationResult = NonNullable<
  Awaited<ReturnType<typeof collectionUpdate>>
>;
export type CollectionUpdateMutationError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useCollectionUpdate = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  collectionMark: string,
  options?: {
    swr?: SWRMutationConfiguration<
      Awaited<ReturnType<typeof collectionUpdate>>,
      TError,
      Key,
      CollectionUpdateBody,
      Awaited<ReturnType<typeof collectionUpdate>>
    > & { swrKey?: string };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey =
    swrOptions?.swrKey ?? getCollectionUpdateMutationKey(collectionMark);
  const swrFn = getCollectionUpdateMutationFetcher(collectionMark);

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
/**
 * Delete a collection owned by the authenticated account.
 */
export const collectionDelete = (collectionMark: string) => {
  return fetcher<void>({
    url: `/collections/${collectionMark}`,
    method: "DELETE",
  });
};

export const getCollectionDeleteMutationFetcher = (collectionMark: string) => {
  return (_: Key, __: { arg: Arguments }): Promise<void> => {
    return collectionDelete(collectionMark);
  };
};
export const getCollectionDeleteMutationKey = (collectionMark: string) =>
  [`/collections/${collectionMark}`] as const;

export type CollectionDeleteMutationResult = NonNullable<
  Awaited<ReturnType<typeof collectionDelete>>
>;
export type CollectionDeleteMutationError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useCollectionDelete = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  collectionMark: string,
  options?: {
    swr?: SWRMutationConfiguration<
      Awaited<ReturnType<typeof collectionDelete>>,
      TError,
      Key,
      Arguments,
      Awaited<ReturnType<typeof collectionDelete>>
    > & { swrKey?: string };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey =
    swrOptions?.swrKey ?? getCollectionDeleteMutationKey(collectionMark);
  const swrFn = getCollectionDeleteMutationFetcher(collectionMark);

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
/**
 * Add a post to a collection. The collection must be owned by the account
making the request. The post can be any published post of any kind.

 */
export const collectionAddPost = (collectionMark: string, postId: string) => {
  return fetcher<CollectionAddPostOKResponse>({
    url: `/collections/${collectionMark}/posts/${postId}`,
    method: "PUT",
  });
};

export const getCollectionAddPostMutationFetcher = (
  collectionMark: string,
  postId: string,
) => {
  return (
    _: Key,
    __: { arg: Arguments },
  ): Promise<CollectionAddPostOKResponse> => {
    return collectionAddPost(collectionMark, postId);
  };
};
export const getCollectionAddPostMutationKey = (
  collectionMark: string,
  postId: string,
) => [`/collections/${collectionMark}/posts/${postId}`] as const;

export type CollectionAddPostMutationResult = NonNullable<
  Awaited<ReturnType<typeof collectionAddPost>>
>;
export type CollectionAddPostMutationError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useCollectionAddPost = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  collectionMark: string,
  postId: string,
  options?: {
    swr?: SWRMutationConfiguration<
      Awaited<ReturnType<typeof collectionAddPost>>,
      TError,
      Key,
      Arguments,
      Awaited<ReturnType<typeof collectionAddPost>>
    > & { swrKey?: string };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey =
    swrOptions?.swrKey ??
    getCollectionAddPostMutationKey(collectionMark, postId);
  const swrFn = getCollectionAddPostMutationFetcher(collectionMark, postId);

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
/**
 * Remove a post from a collection. The collection must be owned by the
account making the request.

 */
export const collectionRemovePost = (
  collectionMark: string,
  postId: string,
) => {
  return fetcher<CollectionRemovePostOKResponse>({
    url: `/collections/${collectionMark}/posts/${postId}`,
    method: "DELETE",
  });
};

export const getCollectionRemovePostMutationFetcher = (
  collectionMark: string,
  postId: string,
) => {
  return (
    _: Key,
    __: { arg: Arguments },
  ): Promise<CollectionRemovePostOKResponse> => {
    return collectionRemovePost(collectionMark, postId);
  };
};
export const getCollectionRemovePostMutationKey = (
  collectionMark: string,
  postId: string,
) => [`/collections/${collectionMark}/posts/${postId}`] as const;

export type CollectionRemovePostMutationResult = NonNullable<
  Awaited<ReturnType<typeof collectionRemovePost>>
>;
export type CollectionRemovePostMutationError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useCollectionRemovePost = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  collectionMark: string,
  postId: string,
  options?: {
    swr?: SWRMutationConfiguration<
      Awaited<ReturnType<typeof collectionRemovePost>>,
      TError,
      Key,
      Arguments,
      Awaited<ReturnType<typeof collectionRemovePost>>
    > & { swrKey?: string };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey =
    swrOptions?.swrKey ??
    getCollectionRemovePostMutationKey(collectionMark, postId);
  const swrFn = getCollectionRemovePostMutationFetcher(collectionMark, postId);

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
/**
 * Add a node to a collection. The collection must be owned by the account
making the request. The node can be any published node or any node
not published but owned by the collection owner.

 */
export const collectionAddNode = (collectionMark: string, nodeId: string) => {
  return fetcher<CollectionAddNodeOKResponse>({
    url: `/collections/${collectionMark}/nodes/${nodeId}`,
    method: "PUT",
  });
};

export const getCollectionAddNodeMutationFetcher = (
  collectionMark: string,
  nodeId: string,
) => {
  return (
    _: Key,
    __: { arg: Arguments },
  ): Promise<CollectionAddNodeOKResponse> => {
    return collectionAddNode(collectionMark, nodeId);
  };
};
export const getCollectionAddNodeMutationKey = (
  collectionMark: string,
  nodeId: string,
) => [`/collections/${collectionMark}/nodes/${nodeId}`] as const;

export type CollectionAddNodeMutationResult = NonNullable<
  Awaited<ReturnType<typeof collectionAddNode>>
>;
export type CollectionAddNodeMutationError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useCollectionAddNode = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  collectionMark: string,
  nodeId: string,
  options?: {
    swr?: SWRMutationConfiguration<
      Awaited<ReturnType<typeof collectionAddNode>>,
      TError,
      Key,
      Arguments,
      Awaited<ReturnType<typeof collectionAddNode>>
    > & { swrKey?: string };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey =
    swrOptions?.swrKey ??
    getCollectionAddNodeMutationKey(collectionMark, nodeId);
  const swrFn = getCollectionAddNodeMutationFetcher(collectionMark, nodeId);

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
/**
 * Remove a node from a collection. The collection must be owned by the
account making the request.

 */
export const collectionRemoveNode = (
  collectionMark: string,
  nodeId: string,
) => {
  return fetcher<CollectionRemoveNodeOKResponse>({
    url: `/collections/${collectionMark}/nodes/${nodeId}`,
    method: "DELETE",
  });
};

export const getCollectionRemoveNodeMutationFetcher = (
  collectionMark: string,
  nodeId: string,
) => {
  return (
    _: Key,
    __: { arg: Arguments },
  ): Promise<CollectionRemoveNodeOKResponse> => {
    return collectionRemoveNode(collectionMark, nodeId);
  };
};
export const getCollectionRemoveNodeMutationKey = (
  collectionMark: string,
  nodeId: string,
) => [`/collections/${collectionMark}/nodes/${nodeId}`] as const;

export type CollectionRemoveNodeMutationResult = NonNullable<
  Awaited<ReturnType<typeof collectionRemoveNode>>
>;
export type CollectionRemoveNodeMutationError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useCollectionRemoveNode = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  collectionMark: string,
  nodeId: string,
  options?: {
    swr?: SWRMutationConfiguration<
      Awaited<ReturnType<typeof collectionRemoveNode>>,
      TError,
      Key,
      Arguments,
      Awaited<ReturnType<typeof collectionRemoveNode>>
    > & { swrKey?: string };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey =
    swrOptions?.swrKey ??
    getCollectionRemoveNodeMutationKey(collectionMark, nodeId);
  const swrFn = getCollectionRemoveNodeMutationFetcher(collectionMark, nodeId);

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
