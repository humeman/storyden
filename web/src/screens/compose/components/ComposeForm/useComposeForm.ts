import { zodResolver } from "@hookform/resolvers/zod";
import { useRouter } from "next/navigation";
import { useForm } from "react-hook-form";
import { z } from "zod";

import { useCategoryList } from "src/api/openapi-client/categories";
import { threadCreate, threadUpdate } from "src/api/openapi-client/threads";
import { Thread, ThreadInitialProps, Visibility } from "src/api/openapi-schema";

import { handle } from "@/api/client";

export type Props = { editing?: string; initialDraft?: Thread };

export const FormShapeSchema = z.object({
  title: z.string().default(""),
  body: z.string().min(1),
  category: z.string(),
  tags: z.string().array().optional(),
  url: z.string().optional(),
});
export type FormShape = z.infer<typeof FormShapeSchema>;

export function useComposeForm({ initialDraft, editing }: Props) {
  const router = useRouter();

  const { data } = useCategoryList();
  const formContext = useForm<FormShape>({
    resolver: zodResolver(FormShapeSchema),
    reValidateMode: "onChange",
    defaultValues: initialDraft
      ? {
          title: initialDraft.title,
          body: initialDraft.body,
          tags: initialDraft.tags.map((t) => t.name),
          url: initialDraft.link?.url,
        }
      : {
          // hack: the underlying category list select component can't do this.
          category: data?.categories[0]?.id,
        },
  });

  const doSave = async (data: FormShape) => {
    const payload: ThreadInitialProps = {
      ...data,

      // When saving a new draft, these are optional but must be explicitly set.
      title: data.title ?? "",
      body: data.body ?? "",
      url: data.url ?? "",
      tags: data.tags ?? [],

      visibility: Visibility.draft,
    };

    if (editing) {
      await threadUpdate(editing, payload);
    } else {
      const { id } = await threadCreate(payload);

      router.push(`/new?id=${id}`);
    }
  };

  const doPublish = async ({ title, body, category, tags, url }: FormShape) => {
    if (title.length < 1) {
      formContext.setError("title", {
        message: "Your post must have a title to be published",
      });
      return;
    }

    if (editing) {
      const { slug } = await threadUpdate(editing, {
        title,
        body,
        category,
        visibility: Visibility.published,
        tags,
        url,
      });
      router.push(`/t/${slug}`);
    } else {
      const { slug } = await threadCreate({
        title,
        body,
        category,
        visibility: Visibility.published,
        tags,
        url,
      });
      router.push(`/t/${slug}`);
    }
  };

  const onAssetUpload = async () => {
    const state = formContext.getValues();
    await handle(async () => {
      await doSave(state);
    });
  };

  const onAssetDelete = async () => {
    const state = formContext.getValues();
    await handle(async () => {
      await doSave(state);
    });
  };

  function onBack() {
    router.back();
  }

  const onSave = formContext.handleSubmit((data) =>
    handle(async () => {
      await doSave(data);
    }),
  );

  const onPublish = formContext.handleSubmit((data) =>
    handle(async () => {
      await doPublish(data);
    }),
  );

  return {
    onBack,
    onSave,
    onPublish,
    onAssetUpload,
    onAssetDelete,
    formContext,
  };
}
