package wallet

import(
	"io"
	"os"
	"log"
	"time"
	"crypto/rsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/SaKu2110/oauth_chain/wallet/model"
	"github.com/SaKu2110/oauth_chain/wallet/config"
)

type Block struct {
	Index			int `json:"index"`
	Timestamp		int64 `json:"timestamp"`
	Transaction		model.Transaction `json:"transaction"`
	PreviousHash	string `json:"hash"`
	Nonce			int `json:"nonce"`
}

type Wallet struct {
	index	int
	previousHash []byte
}

var(
	publicKey *rsa.PublicKey
	rng io.Reader
	label []byte
)

func init() {
	privateKey, err := config.ReadRsaPrivateKey(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
	}
	rng = rand.Reader
	label = []byte("label")
	publicKey = &privateKey.PublicKey
}

func InitializeWallet() (*Wallet, error){
	originBlock := Block{
		Index: 0,
		Timestamp: time.Now().UnixNano(),
	}
	block, err := json.Marshal(originBlock)
	if err != nil {
		return nil, err
	}
	hash := sha256.Sum256(
		[]byte(string(block)),
	)
	wallet := Wallet{
		index: 0,
		previousHash: hash[:],
	}
	return &wallet, nil
}


func (wallet *Wallet)CreateNewTransaction(transaction model.Transaction) ([]byte, error){
	var crypherTransaction []byte
	wallet.index += 1

	newBlock := Block{
		Index: wallet.index,
		Timestamp: time.Now().UnixNano(),
		PreviousHash: hex.EncodeToString(wallet.previousHash[:]),
		Transaction: transaction,
	}
	block, err := json.Marshal(newBlock)
	if err != nil {
		return nil, err
	}

	crypherTransaction, err = rsa.EncryptOAEP(sha256.New(), rng, publicKey, block, label)
    if err != nil {
        return nil, err
    }

	return crypherTransaction, nil
}