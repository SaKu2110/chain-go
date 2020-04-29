# chain_dev
![chain_dev](https://github.com/SaKu2110/chain_dev/blob/master/chain_dev.jpg "chain_dev")
## Setup
### golangのソースコード生成
```
$ protoc --go_out=plugins=grpc:./ ./network.proto
```
### localhostで実行する
```
$ sh build.sh
$ sh launch_master.sh
```
別のターミナルを開いて以下を実行する
```
$ sh launch_miner.sh
```