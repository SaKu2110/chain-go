#!/bin/sh

go build -o master pool/main.go
go build -o miner node/main.go
