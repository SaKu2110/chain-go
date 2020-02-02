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
	nonce			int `json:"nonce"`
	previousHash	string`json:"previoushash"`
}

type BlockChain struct {
	Hash			[]byte
	previousHash	[]byte
	block			Block
}

// mine終了制約
var difficulty = "2110"

func (block *Block)CreateHash() ([]byte, error){
	byteData, err := json.Marshal(block)
	if err != nil {
		// エラー処理
	}
	data := string(byteData)
	hash := sha256.Sum256(
		[]byte(data),
	)
	return hash[:], nil
}

func (blch *BlockChain)CreateNewTransaction(transaction model.Data, timestamp int64){
	newBlock := Block{
		Index: blch.block.Index + 1,
		Timestamp: timestamp,
		Transaction: transaction,
		previousHash: hex.EncodeToString(blch.Hash[:]),
	}
	blch.Block = newBlock
}

func (blch *BlockChain)Mine(nonce int) ([]byte, error){

	blch.block.nonce = nonce
	hash, err := block.CreateHash()
	// 生成したハッシュを検査
	return hash, err
}