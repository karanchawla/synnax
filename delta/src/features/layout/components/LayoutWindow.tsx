import { Space } from "@synnaxlabs/pluto";
import { appWindow } from "@tauri-apps/api/window";

import { useSelectLayout } from "../store";

import { LayoutContent } from "./LayoutContent";

import { NavTop } from "@/components";

import "./LayoutWindow.css";

export const LayoutWindow = (): JSX.Element => {
  const { label: key } = appWindow;
  const layout = useSelectLayout(key);
  const content = <LayoutContent layoutKey={key} />;
  if (layout?.window?.navTop === true)
    return (
      <Space empty className="delta-main">
        <NavTop />
        {content}
      </Space>
    );
  return content;
};
