FROM golang:1.13.10-alpine3.10
WORKDIR /go/src/github.com/SaKu2110/chain_dev
ADD ./ ./
RUN go build -o pool_api ./pool/main.go
RUN go build -o miner_api ./node/main.go