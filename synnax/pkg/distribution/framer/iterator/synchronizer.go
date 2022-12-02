package iterator

import (
	"context"
	"github.com/synnaxlabs/synnax/pkg/distribution/framer/core"
	"github.com/synnaxlabs/x/confluence"
)

type synchronizer struct {
	internal core.Synchronizer
	confluence.LinearTransform[Response, Response]
}

func newSynchronizer(nodeCount int) confluence.Segment[Response, Response] {
	s := &synchronizer{}
	s.internal.NodeCount = nodeCount
	s.internal.SeqNum = 1
	s.LinearTransform.Transform = s.sync
	return s
}

func (a *synchronizer) sync(_ context.Context, res Response) (Response, bool, error) {
	ack, seqNum, fulfilled := a.internal.Sync(res.SeqNum, res.Ack)
	return Response{
		Variant: AckResponse,
		Command: res.Command,
		Ack:     ack,
		SeqNum:  seqNum,
	}, fulfilled, nil
}
