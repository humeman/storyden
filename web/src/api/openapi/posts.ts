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
  InternalServerErrorResponse,
  NotFoundResponse,
  PostCreateBody,
  PostCreateOKResponse,
  PostReactAddBody,
  PostReactAddOKResponse,
  PostSearchOKResponse,
  PostSearchParams,
  PostUpdateBody,
  PostUpdateOKResponse,
  UnauthorisedResponse,
} from "./schemas";

type AwaitedInput<T> = PromiseLike<T> | T;

type Awaited<O> = O extends AwaitedInput<infer T> ? T : never;

/**
 * Create a new post within a thread.
 */
export const postCreate = (
  threadMark: string,
  postCreateBody: PostCreateBody,
) => {
  return fetcher<PostCreateOKResponse>({
    url: `/v1/threads/${threadMark}/posts`,
    method: "post",
    headers: { "Content-Type": "application/json" },
    data: postCreateBody,
  });
};

/**
 * Publish changes to a single post.
 */
export const postUpdate = (postId: string, postUpdateBody: PostUpdateBody) => {
  return fetcher<PostUpdateOKResponse>({
    url: `/v1/posts/${postId}`,
    method: "patch",
    headers: { "Content-Type": "application/json" },
    data: postUpdateBody,
  });
};

/**
 * Archive a post using soft-delete.
 */
export const postDelete = (postId: string) => {
  return fetcher<void>({ url: `/v1/posts/${postId}`, method: "delete" });
};

/**
 * Search through posts using various queries and filters.
 */
export const postSearch = (params?: PostSearchParams) => {
  return fetcher<PostSearchOKResponse>({
    url: `/v1/posts/search`,
    method: "get",
    params,
  });
};

export const getPostSearchKey = (params?: PostSearchParams) =>
  [`/v1/posts/search`, ...(params ? [params] : [])] as const;

export type PostSearchQueryResult = NonNullable<
  Awaited<ReturnType<typeof postSearch>>
>;
export type PostSearchQueryError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const usePostSearch = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  params?: PostSearchParams,
  options?: {
    swr?: SWRConfiguration<Awaited<ReturnType<typeof postSearch>>, TError> & {
      swrKey?: Key;
      enabled?: boolean;
    };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false;
  const swrKey =
    swrOptions?.swrKey ?? (() => (isEnabled ? getPostSearchKey(params) : null));
  const swrFn = () => postSearch(params);

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
 * Add a reaction to a post.
 */
export const postReactAdd = (
  postId: string,
  postReactAddBody: PostReactAddBody,
) => {
  return fetcher<PostReactAddOKResponse>({
    url: `/v1/posts/${postId}/reacts`,
    method: "put",
    headers: { "Content-Type": "application/json" },
    data: postReactAddBody,
  });
};
