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
  CategoryCreateBody,
  CategoryCreateOKResponse,
  CategoryListOKResponse,
  CategoryUpdateOrderBody,
  InternalServerErrorResponse,
} from "./schemas";

type AwaitedInput<T> = PromiseLike<T> | T;

type Awaited<O> = O extends AwaitedInput<infer T> ? T : never;

/**
 * Create a category for organising posts.
 */
export const categoryCreate = (categoryCreateBody: CategoryCreateBody) => {
  return fetcher<CategoryCreateOKResponse>({
    url: `/v1/categories`,
    method: "post",
    headers: { "Content-Type": "application/json" },
    data: categoryCreateBody,
  });
};

/**
 * Get a list of all categories on the site.
 */
export const categoryList = () => {
  return fetcher<CategoryListOKResponse>({
    url: `/v1/categories`,
    method: "get",
  });
};

export const getCategoryListKey = () => [`/v1/categories`] as const;

export type CategoryListQueryResult = NonNullable<
  Awaited<ReturnType<typeof categoryList>>
>;
export type CategoryListQueryError = InternalServerErrorResponse;

export const useCategoryList = <
  TError = InternalServerErrorResponse,
>(options?: {
  swr?: SWRConfiguration<Awaited<ReturnType<typeof categoryList>>, TError> & {
    swrKey?: Key;
    enabled?: boolean;
  };
}) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false;
  const swrKey =
    swrOptions?.swrKey ?? (() => (isEnabled ? getCategoryListKey() : null));
  const swrFn = () => categoryList();

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
 * Update the sort order of categories.
 */
export const categoryUpdateOrder = (
  categoryUpdateOrderBody: CategoryUpdateOrderBody,
) => {
  return fetcher<CategoryListOKResponse>({
    url: `/v1/categories`,
    method: "patch",
    headers: { "Content-Type": "application/json" },
    data: categoryUpdateOrderBody,
  });
};
