import { Synnax, TimeStamp } from "@synnaxlabs/client";

// This example demonstrates how to stream live data from a channel in Synnax.
// Live-streaming is useful for real-time data processing and analysis, and is an
// integral part of Synnax's control sequence and data streaming capabilities.

// This example is meant to be used in conjunction with the stream_write.py example, and
// assumes that example is running in a separate terminal.

// Connect to a locally running, insecure Synnax cluster. If your connection parameters
// are different, enter them here.
const client = new Synnax({
    host: "localhost",
    port: 9090,
    username: "synnax",
    password: "seldon",
    secure: false
});

// We can just specify the names of the channels we'd like to stream from.
const read_from = ["stream_write_example_time", "stream_write_example_data"]

const streamer = await client.telem.newStreamer(read_from);

// It's very important that we close the streamer when we're done with it to release 
// network connections and other resources, so we wrap the streaming loop in a try-finally
// block.
try {
    // Loop through the frames in the streamer. Each iteration will block until a new 
    // frame is available, and then we'll just print it out.
    for await (const frame of streamer) 
        console.log({
            time: new TimeStamp(frame.get("stream_write_example_time")[0].at(0)).toString(),
            data: frame.get("stream_write_example_data")[0].at(0)
        })
} finally {
    streamer.close();
    // Close the client when we're done with it.
    client.close();
}
