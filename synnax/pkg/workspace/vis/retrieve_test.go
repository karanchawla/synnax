// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

package vis_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/synnaxlabs/synnax/pkg/workspace/vis"
)

var _ = Describe("Retrieve", func() {
	It("Should retrieve a Vis", func() {
		p := vis.Vis{Name: "test", Data: "data"}
		Expect(svc.NewWriter(tx).Create(ctx, ws.Key, &p)).To(Succeed())
		var res vis.Vis
		Expect(svc.NewRetrieve().WhereKeys(p.Key).Entry(&res).Exec(ctx, tx)).To(Succeed())
		Expect(res).To(Equal(p))
	})
})
