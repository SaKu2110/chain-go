package main

import(
	"os"
	"fmt"
	"log"
	"net"
	"time"
	"sync"
	"google.golang.org/grpc"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/SaKu2110/chain_dev/chain"
	"github.com/SaKu2110/chain_dev/network"
	"github.com/SaKu2110/chain_dev/pool/config"
	"github.com/SaKu2110/chain_dev/pool/controller"
)

func initializeDataBase() (*gorm.DB, error) {
	var(
		db *gorm.DB
		err error
	)
	token := config.GetConnectionToken()

	count := 1
	for {
		if count > 5 {
			return nil, fmt.Errorf("faild mysql connection")
		}
		db, err = gorm.Open("mysql", token)
		if err != nil {
			return db, nil
		}
		time.Sleep(3 * time.Second)
		count++
	}

	return nil, err
}

func initializeController(db *gorm.DB) (controller.Controller) {
	gbc := chain.Chain{}
	gbc.CreateGenesisBlock()

	return controller.Controller{
		// TODO: 後に廃止する部分
		Chain: &gbc,

		DB: db,
		Nodes: make(map[string]network.NodeNetwork_ShareResultServer),
		Mutex: sync.RWMutex{},
	}
}

func main(){
	var db *gorm.DB
	addr := ":" + os.Getenv("LISTEN_ADDR")
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()

	// db, err = initializeDataBase()
	// if err != nil { log.Fatalf("failed initialize db. err=%s", err) }
	ctrl := initializeController(db)

	network.RegisterNodeNetworkServer(server, &ctrl)
	fmt.Printf("launch server => localhost%s\n", addr)

	err = server.Serve(listen)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}