import { Skeleton, SkeletonText } from "@chakra-ui/react";

import { Unready } from "src/components/site/Unready";

import { VStack } from "@/styled-system/jsx";

import { ComposeForm } from "./components/ComposeForm/ComposeForm";
import { Props, useComposeScreen } from "./useComposeScreen";

export function ComposeScreen(props: Props) {
  const { loadingDraft, draft } = useComposeScreen(props);

  if (loadingDraft)
    return (
      <Unready>
        <Skeleton height={8} />
        <SkeletonText noOfLines={3} />
      </Unready>
    );

  return (
    <VStack alignItems="start" gap="2" w="full" h="full">
      <ComposeForm {...props} initialDraft={draft} />
    </VStack>
  );
}
