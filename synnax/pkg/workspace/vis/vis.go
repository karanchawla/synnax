// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

package vis

import (
	"github.com/google/uuid"
	"github.com/synnaxlabs/x/gorp"
)

type Vis struct {
	Key  uuid.UUID `json:"key" msgpack:"key"`
	Type string    `json:"type" msgpack:"type"`
	Name string    `json:"name" msgpack:"name"`
	Data string    `json:"data" msgpack:"data"`
}

var _ gorp.Entry[uuid.UUID] = Vis{}

// GorpKey implements gorp.Entry.
func (p Vis) GorpKey() uuid.UUID { return p.Key }

// SetOptions implements gorp.Entry.
func (p Vis) SetOptions() []interface{} { return nil }
