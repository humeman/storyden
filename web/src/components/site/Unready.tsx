"use client";

import { PropsWithChildren } from "react";

import { APIError } from "src/api/openapi/schemas";
import { Flex, SkeletonText } from "src/theme/components";

import ErrorBanner from "./ErrorBanner";

export function Unready(props: PropsWithChildren<Partial<APIError>>) {
  if (!props.error) {
    return (
      <Flex
        flexDirection="column"
        width="full"
        justifyContent="center"
        p={4}
        gap={4}
      >
        {props.children ?? (
          <>
            <SkeletonText noOfLines={4} />
            <SkeletonText noOfLines={4} />
            <SkeletonText noOfLines={4} />
          </>
        )}
      </Flex>
    );
  }

  return <ErrorBanner {...props} />;
}
