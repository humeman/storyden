/**
 * Generated by orval v6.24.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import type { CategoryName } from "./categoryName";
import type { CategorySlug } from "./categorySlug";
import type { Metadata } from "./metadata";

export interface CategoryCommonProps {
  admin: boolean;
  colour: string;
  description: string;
  meta?: Metadata;
  name: CategoryName;
  slug: CategorySlug;
  sort: number;
}
