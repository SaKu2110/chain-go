// 初期ブロックの生成を行う

package chain

import(
	"log"
	"encoding/json"
	"crypto/sha256"
)

func (c *Chain) CreateGenesisBlock() {
	originBlock := Block{
		Index: 0,
	}
	block, err := json.Marshal(originBlock)
	if err != nil {
		log.Fatalf("failed to marshal block: %v", err)
	}
	hash := sha256.Sum256(block)

	// TODO: DBにいれる
	c.Data = append(c.Data, hash)
}