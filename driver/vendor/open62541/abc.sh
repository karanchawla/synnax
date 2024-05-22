#
# Copyright 2024 Synnax Labs, Inc.
#
# Use of this software is governed by the Business Source License included in the file
# licenses/BSL.txt.
#
# As of the Change Date specified in that file, in accordance with the Business Source
# License, use of this software will be governed by the Apache License, Version 2.0,
# included in the file licenses/APL.txt.
#

sudo rm -r ./open62541/out
sudo rm -r ./open62541/build
mkdir open62541/build && cd open62541/build
cmake -DCMAKE_BUILD_TYPE=RelWithDebInfo -DUA_NAMESPACE_ZERO=FULL -DCMAKE_INSTALL_PREFIX=../out -DCMAKE_OSX_ARCHITECTURES=x86_64 ..
make
make install