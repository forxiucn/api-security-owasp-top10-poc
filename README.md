# API Security OWASP Top 10 POC

本项目用于 API 安全能力验证与演示，覆盖 OWASP API Security Top 10（2019 & 2023），并提供多语言实现（Python、Java、Go）。通过可运行的示例服务端与客户端，便于安全测试、演示与教学。

## 当前目录结构

```text
api-security-owasp-top10/
├── go/
│   ├── 2019/
│   │   ├── server/
│   │   └── client/
│   ├── 2023/
│   │   ├── server/
│   │   └── client/
│   ├── chain/
│   │   ├── server/
│   │   ├── client/
│   │   │   ├── ordered/
│   │   │   └── unordered/
│   │   └── README.md
│   ├── flow/
│   │   ├── server/
│   │   ├── client/
│   │   └── README.md
│   └── README.md
├── java/
│   ├── 2019/
│   │   ├── server/
│   │   └── client/
│   └── 2023/
│       ├── server/
│       └── client/
└── python/
		├── 2019/
		└── 2023/
```

每个子目录下有更详细的 `README.md`，包含编译和运行示例。

## 快速开始（要点）

- Go 示例：服务端现在只接受 `--port` 和 `--contentPath` 两个参数；客户端的 `--addr` 必须包含协议（`http://` 或 `https://`），并且可以包含 content path（例如 `http://host:port/api`）。
- Java 示例：两个 Java 服务都已将 POM 的 `java.version` 更新为 `21`。

下面给出常用命令示例。

### Go（示例：2019）

编译服务端：
```bash
cd go/2019/server
go mod download
go build -o server .
```

运行服务端（示例）：
```bash
# 默认（端口 5019，无 contentPath）
./server

# 指定端口和 contentPath
./server --port 8080 --contentPath /v1/api
```

### Go（flow示例）

编译服务端：
```bash
cd go/flow/server
go mod download
go build -o flow-server .
```

编译客户端：
```bash
cd go/flow/client
go mod download
go build -o flow-client .
```

运行示例：
```bash
# 启动服务端
./flow-server &

# 运行客户端
./flow-client --addr http://127.0.0.1:5060 --username alice --password secret --to bob --amount 100
```

注意：flow示例实现了基于RSA非对称加密的密码传输机制，提升了API安全性。