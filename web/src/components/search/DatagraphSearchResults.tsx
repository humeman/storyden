import { DatagraphNode, DatagraphSearchResult } from "src/api/openapi/schemas";
import { EmptyState } from "src/components/feed/EmptyState";
import { Heading1 } from "src/theme/components/Heading/Index";

import { FeedItem } from "../feed/common/FeedItem/FeedItem";

import { Box, Flex, LinkOverlay, styled } from "@/styled-system/jsx";

type Props = {
  result: DatagraphSearchResult;
};

export function DatagraphSearchResults({ result }: Props) {
  if (!result.items?.length) {
    return <EmptyState />;
  }

  return (
    <styled.ol width="full" display="flex" flexDirection="column" gap="4">
      {result.items.map((v) => (
        <DatagraphResultItem key={v.id} {...v} />
      ))}
    </styled.ol>
  );
}

export function DatagraphResultItem(props: DatagraphNode) {
  const permalink = buildPermalink(props);

  return (
    <Box position="relative">
      <FeedItem>
        <Flex justifyContent="space-between">
          <Heading1 size="sm">
            <LinkOverlay
              //as={NextLink} // TODO
              href={permalink}
            >
              {props.name}
            </LinkOverlay>
          </Heading1>
        </Flex>

        <styled.p lineClamp={3}>{props.description}</styled.p>
        <styled.p lineClamp={3}>{props.kind}</styled.p>
      </FeedItem>
    </Box>
  );
}

function buildPermalink(d: DatagraphNode): string {
  switch (d.kind) {
    case "thread":
      return `/t/${d.slug}`;
    case "reply":
      return `/t/${d.slug}`;
    case "cluster":
      return `/directory/${d.slug}`;
    case "link":
      return `/l/${d.slug}`;
  }
}
