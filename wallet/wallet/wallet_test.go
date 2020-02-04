package wallet

import(
	"fmt"
	"testing"
	"crypto/rsa"
	"crypto/sha256"
	"github.com/SaKu2110/oauth_chain/wallet/model"
	"github.com/SaKu2110/oauth_chain/wallet/config"
)

func TestCreateTransaction(t *testing.T){
	var block, blockData []byte
	var privateKey *rsa.PrivateKey

	wallet, err := InitializeWallet()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	transaction := model.Transaction{
		SenderAddress: "Bob",
		ReceiverAddress: "Alice",
		Amount: 10.0,
	}
	block, err = wallet.CreateNewTransaction(transaction)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	// 複合
	privateKey, err = config.ReadRsaPrivateKey("/Users/shotaro-yamada/go/src/github.com/SaKu2110/oauth_chain/private_key.pem")
	blockData, err = rsa.DecryptOAEP(sha256.New(), rng, privateKey, block, label)
	fmt.Println(string(blockData))

}