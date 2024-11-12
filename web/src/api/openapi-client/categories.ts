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
import useSWRMutation from "swr/mutation";
import type { SWRMutationConfiguration } from "swr/mutation";

import { fetcher } from "../client";
import type {
  BadRequestResponse,
  CategoryCreateBody,
  CategoryCreateOKResponse,
  CategoryListOKResponse,
  CategoryUpdateBody,
  CategoryUpdateOKResponse,
  CategoryUpdateOrderBody,
  InternalServerErrorResponse,
  UnauthorisedResponse,
} from "../openapi-schema";

/**
 * Create a category for organising posts.
 */
export const categoryCreate = (categoryCreateBody: CategoryCreateBody) => {
  return fetcher<CategoryCreateOKResponse>({
    url: `/categories`,
    method: "POST",
    headers: { "Content-Type": "application/json" },
    data: categoryCreateBody,
  });
};

export const getCategoryCreateMutationFetcher = () => {
  return (
    _: Key,
    { arg }: { arg: CategoryCreateBody },
  ): Promise<CategoryCreateOKResponse> => {
    return categoryCreate(arg);
  };
};
export const getCategoryCreateMutationKey = () => [`/categories`] as const;

export type CategoryCreateMutationResult = NonNullable<
  Awaited<ReturnType<typeof categoryCreate>>
>;
export type CategoryCreateMutationError =
  | BadRequestResponse
  | UnauthorisedResponse
  | InternalServerErrorResponse;

export const useCategoryCreate = <
  TError =
    | BadRequestResponse
    | UnauthorisedResponse
    | InternalServerErrorResponse,
>(options?: {
  swr?: SWRMutationConfiguration<
    Awaited<ReturnType<typeof categoryCreate>>,
    TError,
    Key,
    CategoryCreateBody,
    Awaited<ReturnType<typeof categoryCreate>>
  > & { swrKey?: string };
}) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey = swrOptions?.swrKey ?? getCategoryCreateMutationKey();
  const swrFn = getCategoryCreateMutationFetcher();

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
/**
 * Get a list of all categories on the site.
 */
export const categoryList = () => {
  return fetcher<CategoryListOKResponse>({ url: `/categories`, method: "GET" });
};

export const getCategoryListKey = () => [`/categories`] as const;

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
    url: `/categories`,
    method: "PATCH",
    headers: { "Content-Type": "application/json" },
    data: categoryUpdateOrderBody,
  });
};

export const getCategoryUpdateOrderMutationFetcher = () => {
  return (
    _: Key,
    { arg }: { arg: CategoryUpdateOrderBody },
  ): Promise<CategoryListOKResponse> => {
    return categoryUpdateOrder(arg);
  };
};
export const getCategoryUpdateOrderMutationKey = () => [`/categories`] as const;

export type CategoryUpdateOrderMutationResult = NonNullable<
  Awaited<ReturnType<typeof categoryUpdateOrder>>
>;
export type CategoryUpdateOrderMutationError = InternalServerErrorResponse;

export const useCategoryUpdateOrder = <
  TError = InternalServerErrorResponse,
>(options?: {
  swr?: SWRMutationConfiguration<
    Awaited<ReturnType<typeof categoryUpdateOrder>>,
    TError,
    Key,
    CategoryUpdateOrderBody,
    Awaited<ReturnType<typeof categoryUpdateOrder>>
  > & { swrKey?: string };
}) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey = swrOptions?.swrKey ?? getCategoryUpdateOrderMutationKey();
  const swrFn = getCategoryUpdateOrderMutationFetcher();

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
/**
 * Create a category for organising posts.
 */
export const categoryUpdate = (
  categoryId: string,
  categoryUpdateBody: CategoryUpdateBody,
) => {
  return fetcher<CategoryUpdateOKResponse>({
    url: `/categories/${categoryId}`,
    method: "PATCH",
    headers: { "Content-Type": "application/json" },
    data: categoryUpdateBody,
  });
};

export const getCategoryUpdateMutationFetcher = (categoryId: string) => {
  return (
    _: Key,
    { arg }: { arg: CategoryUpdateBody },
  ): Promise<CategoryUpdateOKResponse> => {
    return categoryUpdate(categoryId, arg);
  };
};
export const getCategoryUpdateMutationKey = (categoryId: string) =>
  [`/categories/${categoryId}`] as const;

export type CategoryUpdateMutationResult = NonNullable<
  Awaited<ReturnType<typeof categoryUpdate>>
>;
export type CategoryUpdateMutationError =
  | BadRequestResponse
  | UnauthorisedResponse
  | InternalServerErrorResponse;

export const useCategoryUpdate = <
  TError =
    | BadRequestResponse
    | UnauthorisedResponse
    | InternalServerErrorResponse,
>(
  categoryId: string,
  options?: {
    swr?: SWRMutationConfiguration<
      Awaited<ReturnType<typeof categoryUpdate>>,
      TError,
      Key,
      CategoryUpdateBody,
      Awaited<ReturnType<typeof categoryUpdate>>
    > & { swrKey?: string };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey = swrOptions?.swrKey ?? getCategoryUpdateMutationKey(categoryId);
  const swrFn = getCategoryUpdateMutationFetcher(categoryId);

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
