#  Copyright 2024 Synnax Labs, Inc.
#
#  Use of this software is governed by the Business Source License included in the file
#  licenses/BSL.txt.
#
#  As of the Change Date specified in that file, in accordance with the Business Source
#  License, use of this software will be governed by the Apache License, Version 2.0,
#  included in the file licenses/APL.txt.

import asyncio
import logging
import datetime
import math
import time

from asyncua import Server, ua
from asyncua.common.methods import uamethod


async def main():
    server = Server()
    await server.init()
    server.set_endpoint("opc.tcp://localhost:4841/freeopcua/server/")
    uri = "http://examples.freeopcua.github.io"
    idx = await server.register_namespace(uri)


    # Populating our address space
    myobj = await server.nodes.objects.add_object(idx, "MyObject")

    # Add different types of variables
    myval = await myobj.add_variable(idx, "my_var_1", 6.7)
    myarray = await myobj.add_variable(idx, "my_array", [1, 2, 3, 4, 5, 6, 7, 8], ua.VariantType.Float)
    my_int_array = await myobj.add_variable(idx, "my_int_array", [1, 2, 3, 4, 5, 6, 7, 8], ua.VariantType.Int32)
    mytimearray = await myobj.add_variable(idx, "my_time_array", [
        datetime.datetime.utcnow(),
        datetime.datetime.utcnow() + datetime.timedelta(milliseconds=1),
        datetime.datetime.utcnow() + datetime.timedelta(milliseconds=2),
        datetime.datetime.utcnow() + datetime.timedelta(milliseconds=3),
        datetime.datetime.utcnow() + datetime.timedelta(milliseconds=4),
        datetime.datetime.utcnow() + datetime.timedelta(milliseconds=5),
        datetime.datetime.utcnow() + datetime.timedelta(milliseconds=6),
        datetime.datetime.utcnow() + datetime.timedelta(milliseconds=7),
        datetime.datetime.utcnow() + datetime.timedelta(milliseconds=8),
    ], ua.VariantType.DateTime)

    myarray.set_writable()
    mytimearray.set_writable()

    RATE = 36*20*12
    ARRAY_SIZE = 40*8
    mytimearray.write_array_dimensions([ARRAY_SIZE])
    myarray.write_array_dimensions([ARRAY_SIZE])

    for i in range(100):
        # add 30 float variables t OPC
        await myobj.add_variable(idx, f"my_float_{i}", i)

    i = 0
    start_ref = datetime.datetime.utcnow()
    async with server:
        while True:
            i += 1
            start = datetime.datetime.utcnow()
            timestamps = [start + datetime.timedelta(seconds=j * ((1 / RATE))) for j in range(ARRAY_SIZE)]
            values = [math.sin((timestamps[j] - start_ref).total_seconds()) for j in range(ARRAY_SIZE)]
            await myarray.set_value(values, varianttype=ua.VariantType.Float)
            await mytimearray.set_value(timestamps, varianttype=ua.VariantType.DateTime)
            duration = (datetime.datetime.utcnow() - start).total_seconds()
            await asyncio.sleep((1/RATE) - duration)


if __name__ == "__main__":
    asyncio.run(main(), debug=True)
