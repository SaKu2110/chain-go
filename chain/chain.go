package chain

type Chain struct {
	Data [][32]byte `json:"data"`	
}

// DBから引っ張ってくる
type Transaction struct {
	Sender	string `json:"sender"`
	Receiver	string `json:"receiver"`
	Timestamp	int64 `json:"timestamp"`
	Amount	float64 `json:"amount"`
}

// Block template
type Block struct {
	Index	int `json:"index"`
	PreviousHash	[]byte `json:"previoushash"`
	Transactions	[]Transaction `json:"transactions"`
	Root	[]byte `json:"root"`
	Nonce	int `json:"nonce"`
}