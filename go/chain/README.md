# Chain scenario (有序/无序调用链)

此目录提供一个调用链示例：一系列按顺序调用的 API（step 1..N）。服务端通过 URL 区分有序与无序两套接口：

- `/chain/ordered/step/:n?chainId=...` — 服务端会校验调用顺序（必须按 1,2,3.. 的顺序执行）。
- `/chain/unordered/step/:n?chainId=...` — 服务端不校验顺序，任意顺序调用均可。

同时提供两个客户端：

- `client/ordered`：按顺序调用 `/chain/ordered/step/1..N`（若中间某步失败会停止）。
- `client/unordered`：以随机顺序调用 `/chain/unordered/step/1..N`。

默认 server 监听端口 `5050`。可以在启动客户端时通过 `--addr` 指定服务地址（必须包含协议）。

编译示例：
```bash
# server
cd go/chain/server
go mod download
go build -o chain-server .

# ordered client
cd ../client/ordered
go mod download
go build -o ordered-client .

# unordered client
cd ../unordered
go mod download
go build -o unordered-client .
```

运行示例：
```bash
# 启动服务
./chain-server

# 先重置 chain 状态（可选）
curl -X POST "http://127.0.0.1:5050/chain/ordered/reset?chainId=my1"

# 有序客户端
./ordered-client --addr http://127.0.0.1:5050 --chainId my1 --steps 5

# 无序客户端
./unordered-client --addr http://127.0.0.1:5050 --chainId my1 --steps 5
```
