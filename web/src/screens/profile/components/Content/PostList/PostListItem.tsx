import { Flex } from "@chakra-ui/react";

import { PostProps } from "src/api/openapi/schemas";
import { ContentViewer } from "src/components/ContentViewer/ContentViewer";
import { Byline } from "src/screens/thread/components/Byline";
import { PostMenu } from "src/screens/thread/components/PostMenu/PostMenu";
import { ReactList } from "src/screens/thread/components/ReactList/ReactList";

export function PostListItem(props: PostProps) {
  return (
    <Flex id={props.id} flexDir="column" gap={2}>
      <ContentViewer value={props.body} />
      <ReactList {...props} />

      <Byline
        href={`#${props.id}`}
        author={props.author.handle}
        time={new Date(props.createdAt)}
        updated={new Date(props.updatedAt)}
        more={<PostMenu {...props} />}
      />
    </Flex>
  );
}
