package controller

import(
	"os"
	"log"
	"sync"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/SaKu2110/chain_dev/chain"
	"github.com/SaKu2110/chain_dev/network"
)

type Controller struct{
	// ここ以下は変更が行われる可能性が高い部分
	Chain *chain.Chain
	// ここ以降は変更がない部分
	network.UnimplementedNodeNetworkServer
	Nodes map[string]network.NodeNetwork_PublishResultServer
	Mutex sync.RWMutex
	DB *gorm.DB
}

func (ctrl *Controller) SyncChain(ctx context.Context, request *network.MinerInfo) (*network.ChainInfo, error) {
	if name := request.GetName(); name == "" {
		log.Printf("pool: Request to share chain information from anonymous user rejected\n")
		return &network.ChainInfo{}, nil
	}
	// blockchain 取得
	list := ctrl.Chain
	// jsonデータ化
	jsonData, err := json.Marshal(&list)
	if err != nil {
		return nil, err
	}
	// TODO: chain暗号化

	log.Printf("pool: Shared chain information to %s\n", request.GetName())
	return &network.ChainInfo{Data: jsonData}, nil
}

func (ctrl *Controller) GetTransaction(ctx context.Context, request *network.MinerInfo) (*network.Transactions, error) {
	if name := request.GetName(); name == "" {
		log.Printf("pool: Request to share transaction from anonymous user rejected\n")
		return &network.Transactions{}, nil
	}
	// transactionの取得
	// TODO: DBから引っ張ってくる
	var trs []chain.Transaction
	// jsonデータ化
	jsonData, err := json.Marshal(&chain.Block{
		Transactions: trs,
	})
	if err != nil {
		return nil, err
	}
	// TODO: 暗号化

	log.Printf("pool: Shared transaction information to %s\n", request.GetName())
	return &network.Transactions{Data: jsonData}, nil
}

func (ctrl *Controller) addClient(userid string, srv network.NodeNetwork_PublishResultServer) {
	ctrl.Mutex.Lock()
	defer ctrl.Mutex.Unlock()
	ctrl.Nodes[userid] = srv
}

func (ctrl *Controller) removeClient(userid string) {
	ctrl.Mutex.Lock()
	defer ctrl.Mutex.Unlock()
	delete(ctrl.Nodes, userid)
}

func (ctrl *Controller) getClients() []network.NodeNetwork_PublishResultServer {
	var cs []network.NodeNetwork_PublishResultServer

	ctrl.Mutex.RLock()
	defer ctrl.Mutex.RUnlock()
	for _, c := range ctrl.Nodes {
		cs = append(cs, c)
	}
	return cs
}

func (ctrl *Controller) PublishResult(srv network.NodeNetwork_PublishResultServer) error {
	userid := uuid.Must(uuid.NewRandom()).String()
	log.Printf("new user: %s", userid)

	ctrl.addClient(userid, srv)
	defer ctrl.removeClient(userid)

	defer func() {
		err := recover()
		if err != nil {
			log.Printf("panic: %v", err)
			os.Exit(1)
		}
	}()
	for {
		recv, err := srv.Recv()
		if err != nil {
			log.Printf("recv err: %v", err)
			break
		}
		// nodeに公開
		for _, ss := range ctrl.getClients() {
			err = ss.Send(&network.MiningInfo{
					Transactions: recv.Transactions,
					Nonce: recv.Nonce,
					Miner: recv.Miner,
				})
			if err != nil {
				log.Printf("broadcast err: %v", err)
			}
		}
	}
	return nil
}

func (ctrl *Controller) ValidateNonce(ctx context.Context, request *network.CheckerInfo) (*network.Response, error){
	switch request.GetStatus() {
	case true:
		log.Printf("node: %s added block\n", request.GetName())
	case false:
		log.Printf("node: %s refused to add block\n", request.GetName())
	}
	return &network.Response{}, nil
}