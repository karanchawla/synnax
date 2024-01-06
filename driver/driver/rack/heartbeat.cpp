
#include "driver/rack/rack.h"

device::Heartbeat::Heartbeat(
        synnax::RackKey rack,
        synnax::Synnax client,
        std::uint32_t generation
) :
        rack_key(rack),
        client(std::move(client)),
        generation(generation),
        version(0),
        running(false),
        exit_err(freighter::NIL) {
}

freighter::Error device::Heartbeat::start() {
    auto [channel, err] = client.channels.retrieve("sy_rack_heartbeat");
    if (err) return err;
    rack_heartbeat_channel = channel;
    running = true;
    exec_thread = std::thread(&Heartbeat::run, this);
    return freighter::NIL;
}

void device::Heartbeat::stop() {
    running = false;
    exec_thread.join();
}

void device::Heartbeat::run() {
    std::vector <synnax::ChannelKey> channels = {rack_heartbeat_channel.key};
    auto [writer, err] = client.telem.openWriter(synnax::WriterConfig{.channels = channels});
    if (err) {
        if (err.type == freighter::TYPE_UNREACHABLE && breaker.wait()) run();
        exit_err = err;
        return;
    }

    while (running) {
        auto heartbeat = static_cast<std::uint64_t>(generation) << 32 | version;
        auto series = synnax::Series(std::vector < std::uint64_t > {heartbeat});
        auto fr = synnax::Frame(1);
        fr.add(rack_heartbeat_channel.key, std::move(series));
        if (!writer.write(std::move(fr))) {
            auto w_err = writer.error();
            if (w_err.type == freighter::TYPE_UNREACHABLE && breaker.wait()) run();
            exit_err = w_err;
            break;
        }
        version++;
    }
    writer.close();
}