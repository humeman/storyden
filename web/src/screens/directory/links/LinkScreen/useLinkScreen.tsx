import { useLinkGet } from "src/api/openapi-client/links";
import { LinkWithRefs } from "src/api/openapi-schema";

export type Props = {
  slug: string;
  link: LinkWithRefs;
};

export function useLinkScreen(props: Props) {
  const { data, error } = useLinkGet(props.slug, {
    swr: { fallbackData: props.link },
  });

  if (!data) {
    return {
      ready: false as const,
      error,
    };
  }

  return {
    ready: true as const,
    data: data,
  };
}
