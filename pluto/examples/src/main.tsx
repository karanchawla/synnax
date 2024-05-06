import ReactDOM from "react-dom/client";

import { Pluto, Canvas } from "@synnaxlabs/pluto";

import { Channel, TimeSpan } from "@synnaxlabs/pluto";

import WorkerURL from "./worker?worker&url";

import "@synnaxlabs/pluto/dist/style.css";

const Main = () => (
  <Pluto.Provider
    workerURL={WorkerURL}
    theming={{ theme: { colors: { primary: "#b57edc" } } }}
    connParams={{
      host: "demo.synnaxlabs.com",
      port: 9090,
      username: "synnax",
      password: "seldon",
      secure: true,
    }}
  >
    <Canvas.Canvas
      style={{
        position: "fixed",
        top: 0,
        left: 0,
        width: "100%",
        height: "100%",
        background: "var(--pluto-gray-l0)",
      }}
    >
      <Channel.LinePlot
        lines={[
          {
            key: "line1",
            axes: { x: "x", y: "y" },
            channels: {
              x: "stream_write_example_time",
              y: "stream_write_example_data_1",
            },
            variant: "dynamic",
            color: "#E87E04",
            label: "Line 1",
            timeSpan: TimeSpan.seconds(30),
            strokeWidth: 3,
          },
        ]}
        axes={[
          {
            key: "x",
            label: "Time",
            location: "bottom",
            color: "#FFFFFF",
            type: "time",
          },
          {
            key: "y",
            label: "Value",
            location: "left",
            color: "#FFFFFF",
          },
        ]}
      />
    </Canvas.Canvas>
  </Pluto.Provider>
);

ReactDOM.createRoot(document.getElementById("root")!).render(<Main />);