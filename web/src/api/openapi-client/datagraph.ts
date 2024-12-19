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
  DatagraphAskOKResponse,
  DatagraphAskParams,
  DatagraphSearchOKResponse,
  DatagraphSearchParams,
  InternalServerErrorResponse,
  NotFoundResponse,
  UnauthorisedResponse,
} from "../openapi-schema";

/**
 * Query and search content.
 */
export const datagraphSearch = (params: DatagraphSearchParams) => {
  return fetcher<DatagraphSearchOKResponse>({
    url: `/datagraph`,
    method: "GET",
    params,
  });
};

export const getDatagraphSearchKey = (params: DatagraphSearchParams) =>
  [`/datagraph`, ...(params ? [params] : [])] as const;

export type DatagraphSearchQueryResult = NonNullable<
  Awaited<ReturnType<typeof datagraphSearch>>
>;
export type DatagraphSearchQueryError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useDatagraphSearch = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  params: DatagraphSearchParams,
  options?: {
    swr?: SWRConfiguration<
      Awaited<ReturnType<typeof datagraphSearch>>,
      TError
    > & { swrKey?: Key; enabled?: boolean };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false;
  const swrKey =
    swrOptions?.swrKey ??
    (() => (isEnabled ? getDatagraphSearchKey(params) : null));
  const swrFn = () => datagraphSearch(params);

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
 * Ask questions about the community's content.
 */
export const datagraphAsk = (params: DatagraphAskParams) => {
  return fetcher<DatagraphAskOKResponse>({
    url: `/datagraph/qna`,
    method: "GET",
    params,
  });
};

export const getDatagraphAskKey = (params: DatagraphAskParams) =>
  [`/datagraph/qna`, ...(params ? [params] : [])] as const;

export type DatagraphAskQueryResult = NonNullable<
  Awaited<ReturnType<typeof datagraphAsk>>
>;
export type DatagraphAskQueryError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useDatagraphAsk = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  params: DatagraphAskParams,
  options?: {
    swr?: SWRConfiguration<Awaited<ReturnType<typeof datagraphAsk>>, TError> & {
      swrKey?: Key;
      enabled?: boolean;
    };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false;
  const swrKey =
    swrOptions?.swrKey ??
    (() => (isEnabled ? getDatagraphAskKey(params) : null));
  const swrFn = () => datagraphAsk(params);

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
