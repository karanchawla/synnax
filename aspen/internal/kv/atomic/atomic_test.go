package atomic_test

import (
	"context"
	. "github.com/onsi/ginkgo/v2"
	"github.com/synnaxlabs/aspen"
	"github.com/synnaxlabs/aspen/internal/cluster/gossip"
	"github.com/synnaxlabs/aspen/internal/node"
	"github.com/synnaxlabs/freighter/fmock"
)

var _ = Describe("OperationSender", func() {
	var (
		net *fmock.Network[gossip.Message, gossip.Message]
	)
	BeforeEach(func() {
		net = fmock.NewNetwork[gossip.Message, gossip.Message]()
	})
	Describe("Send and Receivce", func() {
		var (
			t1           *fmock.UnaryServer[gossip.Message, gossip.Message]
			ctx          context.Context
			cluster      aspen.Cluster
			operationMap map[node.Key][]Operation
		)

		BeforeEach(func() {
			t1 = net.UnaryServer("")
			ctx = context.Background()

		})
	})
})
