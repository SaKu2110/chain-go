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
	network.UnimplementedNodeNetworkServer
	Nodes map[string]network.NodeNetwork_PublishResultServer
	Mutex sync.RWMutex
	DB *gorm.DB
}

func (ctrl *Controller) SyncChain(ctx context.Context, request *network.MinerInfo) (*network.ChainInfo, error) {
	b := []byte("hoge")
	// data複合
	return &network.ChainInfo{Data: b}, nil
}

func (ctrl *Controller) GetTransaction(ctx context.Context, request *network.MinerInfo) (*network.Transactions, error) {
	var trs []chain.Transaction

	// DBから引っ張ってくる

	b, err := json.Marshal(trs)
	if err != nil {
		return nil, err
	}

	return &network.Transactions{Data: b}, nil
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
		_, err := srv.Recv()
		if err != nil {
			log.Printf("recv err: %v", err)
			break
		}
		// データ複合
		// nodeに公開
		for _, ss := range ctrl.getClients() {
			err = ss.Send(&network.Broadcast{
					// data
				})
			if err != nil {
				log.Printf("broadcast err: %v", err)
			}
		}
	}
	return nil
}