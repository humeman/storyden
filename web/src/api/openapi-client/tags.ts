/**
 * Generated by orval v7.2.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
The Storyden API does not adhere to semantic versioning but instead applies a rolling strategy with deprecations and minimal breaking changes. This has been done mainly for a simpler development process and it may be changed to a more fixed versioning strategy in the future. Ultimately, the primary way Storyden tracks versions is dates, there are no set release tags currently.

 * OpenAPI spec version: rolling
 */
import useSwr from "swr";
import type { Key, SWRConfiguration } from "swr";

import { fetcher } from "../client";
import type {
  InternalServerErrorResponse,
  NotFoundResponse,
  TagGetOKResponse,
  TagListOKResponse,
  TagListParams,
} from "../openapi-schema";

/**
 * Get a list of all tags on the site.
 */
export const tagList = (params?: TagListParams) => {
  return fetcher<TagListOKResponse>({ url: `/tags`, method: "GET", params });
};

export const getTagListKey = (params?: TagListParams) =>
  [`/tags`, ...(params ? [params] : [])] as const;

export type TagListQueryResult = NonNullable<
  Awaited<ReturnType<typeof tagList>>
>;
export type TagListQueryError = InternalServerErrorResponse;

export const useTagList = <TError = InternalServerErrorResponse>(
  params?: TagListParams,
  options?: {
    swr?: SWRConfiguration<Awaited<ReturnType<typeof tagList>>, TError> & {
      swrKey?: Key;
      enabled?: boolean;
    };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false;
  const swrKey =
    swrOptions?.swrKey ?? (() => (isEnabled ? getTagListKey(params) : null));
  const swrFn = () => tagList(params);

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
 * Get information about a tag.
 */
export const tagGet = (tagName: string) => {
  return fetcher<TagGetOKResponse>({ url: `/tags/${tagName}`, method: "GET" });
};

export const getTagGetKey = (tagName: string) => [`/tags/${tagName}`] as const;

export type TagGetQueryResult = NonNullable<Awaited<ReturnType<typeof tagGet>>>;
export type TagGetQueryError = NotFoundResponse | InternalServerErrorResponse;

export const useTagGet = <
  TError = NotFoundResponse | InternalServerErrorResponse,
>(
  tagName: string,
  options?: {
    swr?: SWRConfiguration<Awaited<ReturnType<typeof tagGet>>, TError> & {
      swrKey?: Key;
      enabled?: boolean;
    };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false && !!tagName;
  const swrKey =
    swrOptions?.swrKey ?? (() => (isEnabled ? getTagGetKey(tagName) : null));
  const swrFn = () => tagGet(tagName);

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
