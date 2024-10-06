import chroma from "chroma-js";

import { Role } from "@/api/openapi-schema";
import { Badge } from "@/components/ui/badge";

export type Props = {
  role: Role;
};

export function RoleBadge({ role }: Props) {
  const cssVars = badgeColourCSS(role.colour);

  return (
    <Badge
      size="sm"
      style={cssVars}
      bgColor="colorPalette"
      borderColor="colorPalette.muted"
      color="colorPalette.text"
    >
      {role.name}
    </Badge>
  );
}

function badgeColourCSS(c: string) {
  const { bg, bo, fg } = badgeColours(c);

  return {
    "--colors-color-palette-text": fg,
    "--colors-color-palette-muted": bo,
    "--colors-color-palette": bg,
  } as React.CSSProperties;
}

function badgeColours(c: string) {
  const colour = chroma(c);

  const bg = colour.brighten(2).desaturate(3).css();
  const bo = colour.darken(1).desaturate(1).alpha(0.2).css();
  const fg = colour.darken(1).saturate(2).css();

  return { bg, bo, fg };
}
