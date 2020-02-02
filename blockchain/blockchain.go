package blockchain

import(
	"crypto/sha256"
	"encoding/json"
	"encoding/hex"
	"github.com/SaKu2110/oauth_chain/model"
)

type Block struct {
	Index			int `json:"index"`
	Timestamp		int64 `json:"timestamp"`
	Transaction		model.Data `json:"data"`
	Nonce			int `json:"nonce"`
	previousHash	string`json:"previoushash"`
}

type BlockChain struct {
	Hash	[]byte
	block	Block
}

func (block *Block)CreateHash() ([]byte, error){
	byteData, err := json.Marshal(block)
	if err != nil {
	}
	data := string(byteData)
	hash := sha256.Sum256(
		[]byte(data),
	)
	return hash[:], nil
}

func (blch *BlockChain)AddBlock(block Block) error{
	var err error
	block.previousHash = hex.EncodeToString(blch.Hash[:])
	blch.Hash, err = block.CreateHash()
	return err
}