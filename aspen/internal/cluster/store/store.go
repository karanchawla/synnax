// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

// Package store exposes a simple copy-on-read Store for managing cluster state.
// SinkTarget create a new Store, call store.New().
package store

import (
	"context"
	"github.com/google/uuid"
	"github.com/synnaxlabs/aspen/internal/node"
	"github.com/synnaxlabs/x/store"
)

// Store is an interface representing a copy-on-read Store for managing cluster state.
type Store interface {
	// Observable allows the caller to react to state changes. This state is not diffed i.e.
	// any call that modifies the state, even if no actual change occurs, will get sent to the
	// Observable.
	store.Observable[State]
	// ClusterKey returns the cluster key.
	ClusterKey() uuid.UUID
	// SetClusterKey sets the cluster key.
	SetClusterKey(ctx context.Context, key uuid.UUID)
	// SetNode sets a node in state.
	SetNode(context.Context, node.Node)
	// GetNode returns a node from state. Returns false if the node is not found.
	GetNode(node.ID) (node.Node, bool)
	// Merge merges a node.Group into State.Nodes by selecting nodes from group with heartbeats
	// that are either not in State or are older than in State.
	Merge(ctx context.Context, group node.Group)
	// GetHost returns the host node of the Store.
	GetHost() node.Node
	// SetHost sets the host for the Store.
	SetHost(ctx context.Context, node node.Node)
}

func _copy(s State) State {
	return State{Nodes: s.Nodes.Copy(), HostID: s.HostID, ClusterKey: s.ClusterKey}
}

// shouldNotify decides whether we should notify observers
// of the cluster state change. We only notify if:
//
//  1. The cluster key has been set.
//  2. The host node has been set.
//  3. A node has been added or removed from the cluster.
//  4. The state of a node has changed.
//
// We DO NOT notify on heartbeat increments.
func shouldNotify(prevState, nextState State) bool {
	if prevState.ClusterKey != nextState.ClusterKey {
		return false
	}
	if nextState.HostID == 0 {
		return false
	}
	if len(prevState.Nodes) != len(nextState.Nodes) {
		return true
	}
	for id, n := range nextState.Nodes {
		pn, ok := prevState.Nodes[id]
		if !ok || pn.State != n.State {
			return true
		}
	}
	return false
}

// New opens a new empty, invalid Store.
func New(ctx context.Context) Store {
	c := &core{
		Observable: store.ObservableWrap[State](
			store.New(_copy),
			store.ObservableConfig[State]{ShouldNotify: shouldNotify},
		),
	}
	c.Observable.SetState(ctx, State{Nodes: make(node.Group)})
	return c
}

// State is the current state of the cluster as viewed from the host.
type State struct {
	ClusterKey uuid.UUID
	HostID     node.ID
	Nodes      node.Group
}

func (s *State) IsZero() bool {
	return s.ClusterKey == uuid.Nil && s.HostID == 0 && len(s.Nodes) == 0
}

type core struct {
	store.Observable[State]
}

// ClusterKey implements Store.
func (c *core) ClusterKey() uuid.UUID { return c.Observable.PeekState().ClusterKey }

// SetClusterKey implements Store.
func (c *core) SetClusterKey(ctx context.Context, key uuid.UUID) {
	s := c.Observable.PeekState()
	s.ClusterKey = key
	c.Observable.SetState(ctx, s)
}

// GetNode implements Store.
func (c *core) GetNode(id node.ID) (node.Node, bool) {
	n, ok := c.Observable.PeekState().Nodes[id]
	return n, ok
}

// GetHost implements Store.
func (c *core) GetHost() node.Node {
	n, _ := c.GetNode(c.Observable.PeekState().HostID)
	return n
}

// SetHost implements Store.
func (c *core) SetHost(ctx context.Context, n node.Node) {
	snap := c.Observable.CopyState()
	snap.Nodes[n.ID] = n
	snap.HostID = n.ID
	c.Observable.SetState(ctx, snap)
}

// SetNode implements Store.
func (c *core) SetNode(ctx context.Context, n node.Node) {
	snap := c.Observable.CopyState()
	snap.Nodes[n.ID] = n
	c.Observable.SetState(ctx, snap)
}

// Merge implements Store.
func (c *core) Merge(ctx context.Context, other node.Group) {
	snap := c.Observable.CopyState()
	for _, n := range other {
		in, ok := snap.Nodes[n.ID]
		if !ok || n.Heartbeat.OlderThan(in.Heartbeat) {
			snap.Nodes[n.ID] = n
		}
	}
	c.Observable.SetState(ctx, snap)
}
