# OWASP API Top 10 2023 Server

演示 OWASP API Top 10 2023 安全漏洞的服务端应用，基于 Go + Gin 框架实现。

## 前置要求

- Go 1.16 或更高版本
- 需要安装 `github.com/gin-gonic/gin` 依赖

## 编译

编译为可执行文件，支持单架构或交叉编译。

### 单架构编译

```bash
go build -o server .
```

### 交叉编译示例

```bash
# macOS Apple Silicon（ARM64）
GOOS=darwin GOARCH=arm64 go build -o 2023_server_mac_arm64 .

# macOS Intel（amd64）
GOOS=darwin GOARCH=amd64 go build -o 2023_server_mac_amd64 .

# Linux
GOOS=linux GOARCH=amd64 go build -o 2023_server_linux_amd64 .

# Windows
GOOS=windows GOARCH=amd64 go build -o 2023_server_amd64.exe .
```

## 运行

编译后运行二进制文件，支持以下参数：

```bash
./server [OPTIONS]
```

### 参数说明

| 参数 | 默认值 | 说明 | 示例 |
|------|--------|------|------|
| `--port` | 5023 | 监听端口 | `--port 8080` |
| `--contentPath` | 空 | API 路由前缀（content path），例如 `/v1/api` | `--contentPath /v1/api` |

### 运行示例

```bash
# 默认运行（端口 5023，无 content path）
./server

# 指定端口
./server --port 8080

# 指定 content path
./server --contentPath /v1/api

# 组合参数
./server --port 9090 --contentPath /api/v2023
```

<!-- Gin 运行模式由应用默认设置（DebugMode）；如果需要可在运行时修改代码以支持额外的运行参数 -->

## API 端点

服务提供以下 OWASP API Top 10 2023 相关的演示端点：

| 编号 | 类型 | 路径 | 描述 |
|------|------|------|------|
| 1 | GET | `/api1/items/:itemId` | 对象级授权绕过 (BOLA) |
| 2 | POST | `/api2/login` | 认证绕过 |
| 3 | GET | `/api3/userinfo` | 对象属性级授权绕过 (BOPLA) |
| 4 | GET | `/api4/nolimit` | 资源消耗无限制 |
| 5 | GET | `/api5/admin` | 函数级授权绕过 |
| 6 | POST | `/api6/transfer` | 敏感业务流程无限制 |
| 7 | POST | `/api7/ssrf` | 服务端请求伪造 (SSRF) |
| 8 | GET | `/api8/debug` | 安全配置不当 |
| 9 | GET | `/api9/old-api` | 库存管理不当 |
| 10 | POST | `/api10/unsafe` | 不安全的 API 使用 |

## 示例请求

```bash
# 获取物品信息
curl http://127.0.0.1:5023/api1/items/123

# 登录
curl -X POST http://127.0.0.1:5023/api2/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"123456"}'

# 获取用户信息
curl http://127.0.0.1:5023/api3/userinfo

# SSRF 测试
curl -X POST http://127.0.0.1:5023/api7/ssrf \
  -H "Content-Type: application/json" \
  -d '{"url":"http://example.com"}'
```

## 与 2019 版本的差异

OWASP API Top 10 2023 相比 2019 版本的主要更新：

- 第 3 项：从「过度数据暴露」升级为「对象属性级授权绕过 (BOPLA)」
- 第 6 项：从「批量赋值」升级为「敏感业务流程无限制」
- 第 7 项：从「安全配置不当」升级为「SSRF」
- 新增更多实际攻击场景

## 依赖

- `github.com/gin-gonic/gin` - Web 框架

安装依赖：
```bash
go mod download
```

## 许可证

此项目用于教育和安全研究目的。
