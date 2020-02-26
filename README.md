# chain_dev
## Setup
### golangのソースコード生成
```
protoc --go_out=plugins=grpc:./ ./network.proto
```
### 公開鍵・秘密鍵の生成
```
sh key_generator.sh
```
