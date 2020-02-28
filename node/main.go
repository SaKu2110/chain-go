package main

import(
	"log"
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"github.com/SaKu2110/chain_dev/network"
	"github.com/SaKu2110/chain_dev/node/config"
	"github.com/SaKu2110/chain_dev/node/controller"
)

func connectPoolServer(addr string) (*grpc.ClientConn, error) {
	connection, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	return connection, nil
}

func initializeController(connection *grpc.ClientConn) *controller.Controller {
	return &controller.Controller{
		Name: config.GetNodeName(),
		Client: network.NewNodeNetworkClient(connection),
		CTX: context.Background(),
	}
}

func main() {
	connection, err := connectPoolServer("localhost:50051")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer connection.Close()

	miner := initializeController(connection)

	miner.SyncChain()

	var stream network.NodeNetwork_ShareResultClient
	condition := true
	stream, err = miner.ShareResult()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			if !condition {
				continue
			}
			miner.GetTransaction()
	
			// mining
			miner.Block.Mine(miner.Chain)

			jsonData, err := json.Marshal(&miner.Block.Transactions)
			if err != nil {
				log.Fatalf("couldn't change json data: %v", err)
			}
			err = stream.SendMsg(&network.MiningInfo{
				Index: int64(len(miner.Chain) + 1),
				Transactions: jsonData,
				Nonce: int64(miner.Block.Nonce),
				Miner: miner.Name,
			})
			if err != nil {
				log.Fatal(err)
			}
			condition = false
		}
	}()
	for {
		response, err := stream.Recv()
		if err != nil {
			log.Fatal(err)
		}
		// mining 結果を判定
		if response.Miner == miner.Name {
			condition = true
		}


		result, hash := miner.Block.ValidateBlocks(miner.Chain, int(response.Nonce))
		if result {
			miner.Chain = append(miner.Chain, hash)
		}

		miner.ValidateNonce(result)
	}
}