/**
 * Generated by orval v6.24.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import type { Arguments } from "swr";
import useSWRMutation from "swr/mutation";
import type { SWRMutationConfiguration } from "swr/mutation";

import { fetcher } from "../client";

import type {
  AdminSettingsUpdateBody,
  AdminSettingsUpdateOKResponse,
  BadRequestResponse,
  InternalServerErrorResponse,
  UnauthorisedResponse,
} from "./schemas";

/**
 * Update non-env configuration settings for installation.
 */
export const adminSettingsUpdate = (
  adminSettingsUpdateBody: AdminSettingsUpdateBody,
) => {
  return fetcher<AdminSettingsUpdateOKResponse>({
    url: `/v1/admin`,
    method: "PATCH",
    headers: { "Content-Type": "application/json" },
    data: adminSettingsUpdateBody,
  });
};

export const getAdminSettingsUpdateMutationFetcher = () => {
  return (
    _: string,
    { arg }: { arg: Arguments },
  ): Promise<AdminSettingsUpdateOKResponse> => {
    return adminSettingsUpdate(arg as AdminSettingsUpdateBody);
  };
};
export const getAdminSettingsUpdateMutationKey = () => `/v1/admin` as const;

export type AdminSettingsUpdateMutationResult = NonNullable<
  Awaited<ReturnType<typeof adminSettingsUpdate>>
>;
export type AdminSettingsUpdateMutationError =
  | BadRequestResponse
  | UnauthorisedResponse
  | InternalServerErrorResponse;

export const useAdminSettingsUpdate = <
  TError =
    | BadRequestResponse
    | UnauthorisedResponse
    | InternalServerErrorResponse,
>(options?: {
  swr?: SWRMutationConfiguration<
    Awaited<ReturnType<typeof adminSettingsUpdate>>,
    TError,
    string,
    Arguments,
    Awaited<ReturnType<typeof adminSettingsUpdate>>
  > & { swrKey?: string };
}) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey = swrOptions?.swrKey ?? getAdminSettingsUpdateMutationKey();
  const swrFn = getAdminSettingsUpdateMutationFetcher();

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
