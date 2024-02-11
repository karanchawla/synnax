import { Layout } from "@/layout";
import { ReactElement } from "react";
import {
  Button,
  Header,
  Input,
  Nav,
  Align,
  componentRenderProp,
  Status,
} from "@synnaxlabs/pluto";
import { Icon } from "@synnaxlabs/media";
import {checkUpdate, installUpdate } from '@tauri-apps/api/updater';



export const versionWindowLayout: Layout.LayoutState = {
  key: "updateAvailable",
  windowKey: "updateAvailable",
  type: "updateAvailable",
  name: "Skee Yee",
  location: "window",
  window: {
    resizable: false,
    size: { height: 430, width: 650 },
    navTop: true,
    transparent: true,
  },
};
const handleClick = async ()=> {
  const u = await checkUpdate();
  console.log(u)
  // await installUpdate(); 
};



export const UpdateDialog = ({ onClose }: Layout.RendererProps): ReactElement => {

  return (
    <Align.Space>

    <Header.Header level="h4">
      <Header.Title startIcon={<Icon.Cluster />}>Update Available</Header.Title>
    </Header.Header>

    <Nav.Bar location="bottom" size={48}>
      <Nav.Bar.End>

        <Button.Button onClick={handleClick}>
          Update
        </Button.Button>


      </Nav.Bar.End>
    </Nav.Bar>
  </Align.Space>      
  );
};
