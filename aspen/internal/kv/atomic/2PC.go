package atomic

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/synnaxlabs/aspen"
	"github.com/synnaxlabs/aspen/internal/node"
	"github.com/synnaxlabs/freighter"
	"github.com/synnaxlabs/x/kv"
)

type ActionType bool

const (
	Set    ActionType = true
	Delete ActionType = false
)

type Operation struct {
	Key    []byte
	Value  []byte
	Action ActionType
}

type TransactionStatus int

const (
	ServerPrepare TransactionStatus = iota
	ServerCommit
	ServerAbort
	ClientAcknowledge
	ClientAbort
)

type transaction2pc struct {
	transKey    uuid.UUID
	messageType TransactionStatus
	operations  []Operation
}

type txCoordinator struct {
	cluster aspen.Cluster
	client  freighter.UnaryClient[transaction2pc, transaction2pc]
}

func (t *txCoordinator) serverExecute(
	ctx context.Context,
	cluster aspen.Cluster,
	client freighter.UnaryClient[transaction2pc, transaction2pc],
	operationMap map[node.Key][]Operation) (bool, error) {

	txKey := uuid.New()
	messageType := ServerCommit
	// Send the Prepare message to all the nodes
	for key, operations := range operationMap {
		addr, err := cluster.Resolve(key)
		if err != nil {
			return false, err
		}

		retStruct, err := client.Send(ctx, addr, transaction2pc{
			txKey,
			ServerPrepare,
			operations,
		})

		// System should abort if it can't send message or node responds w/ abort
		if retStruct.messageType == ClientAbort || err != nil {
			messageType = ServerAbort
		}
	}

	var err error
	for key, _ := range operationMap {
		addr, err := cluster.Resolve(key)
		if err != nil {
			return false, err
		}

		_, err = client.Send(ctx, addr, transaction2pc{
			transKey:    txKey,
			messageType: messageType,
		})
	}

	return true, err
}

type peerTransactionExecutor struct {
	txns map[uuid.UUID]kv.Tx
	kv   kv.DB
}

func (p *peerTransactionExecutor) handle(ctx context.Context, RQ transaction2pc) (transaction2pc, error) {

	var err error

	if RQ.messageType == ServerPrepare {
		tx := p.kv.OpenTx()
		p.txns[RQ.transKey] = tx

		for _, operation := range RQ.operations {
			key := operation.Key
			value := operation.Value
			action := operation.Action

			if action { // Set
				err = tx.Set(ctx, key, value)
			} else { // Delete
				err = tx.Delete(ctx, key)
			}

			if err != nil {
				return transaction2pc{
					messageType: ClientAbort,
				}, err
			}
		}

		return transaction2pc{
			messageType: ClientAcknowledge,
		}, err
	}

	if RQ.messageType == ServerCommit {
		err = p.txns[RQ.transKey].Commit(ctx)
		delete(p.txns, RQ.transKey)
	} else if RQ.messageType == ServerAbort {
		err = p.txns[RQ.transKey].Close()
	}

	if err != nil {
		return transaction2pc{}, err
	}
	return transaction2pc{}, errors.New("invalid command")
}

func peerListen(
	ctx context.Context,
	cluster aspen.Cluster,
	db kv.DB,
	server freighter.UnaryServer[transaction2pc, transaction2pc]) {

	p := &peerTransactionExecutor{
		txns: make(map[uuid.UUID]kv.Tx),
		kv:   db,
	}
	server.BindHandler(p.handle)
}
