# API Security OWASP Top 10 POC

本项目用于 API 安全能力验证与演示，覆盖 OWASP API Security Top 10（2019 & 2023），并提供多语言实现（Python、Java、Go）。通过可运行的示例服务端与客户端，便于安全测试、演示与教学。

## 当前目录结构

```text
api-security-owasp-top10/
├── go/
│   ├── 2019/
│   │   ├── server/
│   │   └── client/
│   └── 2023/
│       ├── server/
│       └── client/
├── java/
│   ├── 2019/
│   │   └── server/    # 包含 pom.xml 和 Maven wrapper
│   └── 2023/
│       └── server/    # 包含 pom.xml 和 Maven wrapper
└── python/
		├── 2019/
		└── 2023/
```

每个子目录下有更详细的 `README.md`，包含编译和运行示例。

## 快速开始（要点）

- Go 示例：服务端现在只接受 `--port` 和 `--contentPath` 两个参数；客户端的 `--addr` 必须包含协议（`http://` 或 `https://`），并且可以包含 content path（例如 `http://host:port/api`）。
- Java 示例：两个 Java 服务都已将 POM 的 `java.version` 更新为 `21`；仓库中提供了 Maven wrapper（位于各 `java/*/server` 目录），当系统未安装 `mvn` 时可以使用 wrapper 来构建。

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

编译客户端并运行：
```bash
cd ../client
go build -o client .
# 客户端必须使用完整 URL（包含协议），可包含 contentPath
./client --addr http://127.0.0.1:8080/v1/api
```

同样的用法适用于 `go/2023` 子目录（默认端口为 5023，可在各子 README 查看）。

### Java（示例：2019 server）

构建（使用项目内 Maven wrapper，若系统有 `mvn` 也可直接使用）：
```bash
cd java/2019/server
# 第一次运行 wrapper 会从网络下载 Maven 二进制
./mvnw -f pom.xml -B clean package
```

运行（根据项目包和打包方式可能有所不同，参见子目录 README）

注意：POM 已将 `java.version` 设置为 `21`。运行时请确认系统已安装 JDK 21；同时 Spring Boot 版本未在此步骤升级，若运行出现兼容性问题，可能需要升级 Spring Boot 及相关依赖。

## 关键变更摘要

- Go 服务端：`--env` 参数已移除，现只接受 `--port`、`--contentPath`。
- Go 客户端：`--addr` 现在必须包含协议（例如 `http://127.0.0.1:5019`），客户端直接将该值作为 base URL。
- Java：POM 已修改为使用 Java 21，并为项目添加了 Maven wrapper（位于各 Java server 子目录），方便在没有系统 Maven 的环境下构建。

## 进一步操作

- 查看特定实现的详细说明，请进入对应子目录并阅读 `README.md`。
- 需要我为你：
	- 在 CI 中加入构建脚本（Go/Java）？
	- 升级 Java Spring Boot 以确保与 Java 21 完整兼容？
	- 运行端到端测试（启动某一服务并用 client 调用）并将输出返回？

---
如有问题或建议，欢迎创建 issue 或提交 PR。
