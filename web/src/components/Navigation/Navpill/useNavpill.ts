import { debounce } from "lodash";
import { useRouter } from "next/router";
import { ChangeEvent, useEffect, useState } from "react";
import { postSearch } from "src/api/openapi/posts";
import { PostProps } from "src/api/openapi/schemas";
import { useAuthProvider } from "src/auth/useAuthProvider";

export function useNavpill() {
  const { asPath } = useRouter();
  const [isExpanded, setExpanded] = useState(true);
  const { account } = useAuthProvider();
  const [searchQuery, setSearchQuery] = useState("");
  const [searchResults, setSearchResults] = useState<PostProps[]>([]);

  // Close the menu for either navigation events or outside clicks/taps:

  useEffect(() => setExpanded(false), [asPath]);

  function onExpand() {
    console.log("onExpand");
    setExpanded(true);
  }

  function onClose() {
    console.log("onClose");
    setExpanded(false);
  }

  const doSearch = debounce(async (v: string) => {
    postSearch({ body: v })
      .then((results) => setSearchResults(results.results))
      .catch((e) => {
        console.log({ e });
      });
  }, 250);

  async function onSearch(e: ChangeEvent<HTMLInputElement>) {
    const query = e.target.value;

    setSearchQuery(query);

    if (query === "") {
      setSearchResults([]);
      return;
    }

    await doSearch(e.target.value);
  }

  return {
    isExpanded,
    onExpand,
    onClose,
    account,
    searchQuery,
    onSearch,
    searchResults,
  };
}
