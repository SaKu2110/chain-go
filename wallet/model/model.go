package model

import(
	// "encoding/json"
)

type Transaction struct {
	SenderAddress	string `json:"sender"`
	ReceiverAddress	string `json:"receiver"`
	Amount	float64 `json:"amount"`
}