package controller

import(
	"log"
	"context"
	"encoding/json"
	"github.com/SaKu2110/chain_dev/chain"
	"github.com/SaKu2110/chain_dev/network"
)

type Controller struct {
	Name string
	Client	network.NodeNetworkClient
	CTX	context.Context
	Block chain.Block
	Chain [][32]byte
}

func (ctrl *Controller) SyncChain() {
	result, err := ctrl.Client.SyncChain(
		ctrl.CTX,
		&network.MinerInfo{
			Name: ctrl.Name,
		},
	)
	if err != nil {
		log.Fatalf("node: couldn't get chain data: %v", err)
	}
	var chainData chain.Chain
	err = json.Unmarshal(result.GetData(), &chainData)
	if err != nil {
		log.Fatalf("node: JSON data restore failed: %v", err)
	}
	ctrl.Chain = chainData.Data
	log.Printf("node: Completed synchronization with pool server\n")
}

func (ctrl *Controller) GetTransaction() {
	result, err := ctrl.Client.GetTransaction(
		ctrl.CTX,
		&network.MinerInfo{
			Name: ctrl.Name,
		},
	)
	if err != nil {
		log.Fatalf("node: couldn't get transaction: %v", err)
	}
	err = json.Unmarshal(result.GetData(), &ctrl.Block)
	if err != nil {
		log.Fatalf("node: JSON data restore failed: %v", err)
	}
	log.Printf("node: Successfully obtained Transaction\n")
}

// network.ShareReult()を行うためのチャンネルに入る処理
func (ctrl *Controller) ShareResult() (network.NodeNetwork_ShareResultClient, error) {
	stream, err :=  ctrl.Client.ShareResult(ctrl.CTX)
	if err != nil {
		return nil, err
	}
	return stream, nil
}

func (ctrl *Controller) ValidateNonce(status bool) () {
	_, err := ctrl.Client.ValidateNonce(
		ctrl.CTX,
		&network.CheckerInfo{
			Name: ctrl.Name,
			Status: status,
		},
	)
	if err != nil {
		log.Fatalf("node: couldn't send check data: %v", err)
	}
}