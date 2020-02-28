package chain

import(
	"log"
	"encoding/json"
	"crypto/sha256"
)

func createHash(block *Block) ([32]byte, error){
	blockJSON, err := json.Marshal(block)
	if err != nil {
		var n [32]byte
		return n, err
	}
	hash := sha256.Sum256(
		blockJSON,
	)
	return hash, nil
}

func (blck *Block) setBlockData(chain [][32]byte) *Block {
	return &Block{
		Index: len(blck.Transactions) + 1,
		PreviousHash: chain[len(blck.Transactions)][:],
		Transactions: blck.Transactions,
		Root: chain[0][:],
	}
}

func checkHash(hash [32]byte, difficulty int) bool {
	for i := 0; i < difficulty; i++ {
		if hash[i] != 0x00 {return false}
	}
	return true
}

func (blck *Block) Mine(chain [][32]byte) {
	flag := true
	nonce := 0

	block := blck.setBlockData(chain)

	for flag {
		block.Nonce = nonce
		hash, err := createHash(block)
		if err != nil {
			log.Fatalf("couldn't create hash: %v", err)
		}
		if checkHash(hash, 2) {
			blck.Nonce = nonce
			flag = false
		}
		nonce ++
	}
}

func (blck *Block) ValidateBlocks(chain [][32]byte, nonce int) (bool, [32]byte) {
	block := blck.setBlockData(chain)
	block.Nonce = nonce
	hash, err := createHash(block)
	if err != nil {
		log.Fatalf("couldn't create hash: %v", err)
	}

	return checkHash(hash, 2), hash
}