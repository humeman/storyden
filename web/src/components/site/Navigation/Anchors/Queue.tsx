import { QueueListIcon } from "@heroicons/react/24/outline";

import { LinkButtonStyleProps } from "@/components/ui/link-button";

import { Anchor, MenuItem } from "./Anchor";

export const QueueID = "queue";
export const QueueRoute = "/queue";
export const QueueLabel = "Queue";
export const QueueIcon = <QueueListIcon />;

export function QueueAnchor(props: LinkButtonStyleProps) {
  return (
    <Anchor
      id={QueueID}
      route={QueueRoute}
      label={QueueLabel}
      icon={QueueIcon}
      {...props}
    />
  );
}

export function QueueMenuItem() {
  return (
    <MenuItem
      id={QueueID}
      route={QueueRoute}
      label={QueueLabel}
      icon={QueueIcon}
    />
  );
}
