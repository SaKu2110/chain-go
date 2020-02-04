package model

import(
	// "encoding/json"
)

type Transaction struct {
	SenderAddress	string `json:"sender"`
	ReceiverAddress	string `json:"receiver"`
	Amount	float64 `json:"amount"`
}

type Hash struct {
	Height	int
	Hash	string
	Mind	int
	Miner	string
	Size	int
}

type NetWork struct {
	Transaction	[]byte
}