package main

import(
	"os"
	"log"
	"time"
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
	ip := os.Getenv("LISTEN_IP")
	addr := os.Getenv("LISTEN_ADDR")
	connection, err := connectPoolServer(ip + ":" + addr)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer connection.Close()

	miner := initializeController(connection)

	miner.SyncChain()

	var stream network.NodeNetwork_ShareResultClient
	stream, err = miner.ShareResult()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		// Nunce採掘作業
		for {
			// Transactionが空かを判別
			if !miner.GetTransaction() {
				// リクエスト数を減らしてpoolの負荷を下げる
				time.Sleep(time.Second / 2)
				continue
			}
			log.Printf("node: Successfully obtained Transaction\n")
	
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
		}
	}()
	// hash検証作業
	for {
		response, err := stream.Recv()
		if err != nil {
			log.Fatal(err)
		}
		
		// mining 結果を判定
		result, hash := miner.Block.ValidateBlocks(miner.Chain, int(response.Nonce))
		if result {
			miner.Chain = append(miner.Chain, hash)
		}

		miner.ValidateNonce(result)
	}
}