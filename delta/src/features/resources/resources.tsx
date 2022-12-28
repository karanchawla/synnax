import { ReactElement } from "react";

import { OntologyID } from "@synnaxlabs/client";
import type { OntologyResourceType } from "@synnaxlabs/client";
import { AiFillDatabase } from "react-icons/ai";
import { MdOutlineDeviceHub, MdSensors } from "react-icons/md";

import { ClusterIcon } from "@/features/cluster";
import { LayoutPlacer } from "@/features/layout";
import { LinePlotVisualization, createVisualization } from "@/features/visualization";
import { WorkspaceState } from "@/features/workspace";

export interface SelectionContext {
  id: OntologyID;
  placer: LayoutPlacer;
  workspace: WorkspaceState;
}

export interface ResourceType {
  type: OntologyResourceType;
  icon: ReactElement;
  onSelect?: (ctx: SelectionContext) => void;
  hasChildren: boolean;
}

export const resourceTypes: Record<string, ResourceType> = {
  builtin: {
    type: "builtin",
    icon: <AiFillDatabase />,
    hasChildren: true,
  },
  cluster: {
    type: "cluster",
    icon: <ClusterIcon />,
    hasChildren: true,
  },
  node: {
    type: "node",
    icon: <MdOutlineDeviceHub />,
    hasChildren: true,
  },
  channel: {
    type: "channel",
    icon: <MdSensors />,
    hasChildren: false,
    onSelect: ({ placer, id, workspace }: SelectionContext) => {
      placer(
        createVisualization<LinePlotVisualization>({
          channels: {
            y1: [id.key],
            y2: [],
            y3: [],
            y4: [],
            x1: "",
            x2: "",
          },
          ranges: {
            x1: workspace.activeRange != null ? [workspace.activeRange] : [],
            x2: [],
          },
        })
      );
    },
  },
};
