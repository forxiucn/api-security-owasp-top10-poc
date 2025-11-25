# OWASP API Top 10 2019 Server

演示 OWASP API Top 10 2019 安全漏洞的服务端应用，基于 Go + Gin 框架实现。

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
GOOS=darwin GOARCH=arm64 go build -o 2019_server_mac_arm64 .

# macOS Intel（amd64）
GOOS=darwin GOARCH=amd64 go build -o 2019_server_mac_amd64 .

# Linux
GOOS=linux GOARCH=amd64 go build -o 2019_server_linux_amd64 .

# Windows
GOOS=windows GOARCH=amd64 go build -o 2019_server_amd64.exe .
```

## 运行

编译后运行二进制文件，支持以下参数：

```bash
./server [OPTIONS]
```

### 参数说明

| 参数 | 默认值 | 说明 | 示例 |
|------|--------|------|------|
| `--port` | 5019 | 监听端口 | `--port 8080` |
| `--contentPath` | 空 | API 路由前缀（content path），例如 `/v1/api` | `--contentPath /v1/api` |

### 运行示例

```bash
# 默认运行（端口 5019，无 content path）
./server

# 指定端口
./server --port 8080

# 指定 content path
./server --contentPath /v1/api

# 组合参数
./server --port 9090 --contentPath /api/v2019
```

<!-- Gin 运行模式由应用默认设置（DebugMode）；如果需要可在运行时修改代码以支持额外的运行参数 -->

## API 端点

服务提供以下 OWASP API Top 10 相关的演示端点：

| 编号 | 类型 | 路径 | 描述 |
|------|------|------|------|
| 1 | GET | `/api1/items/:itemId` | 对象级授权绕过 (BOLA) |
| 2 | POST | `/api2/login` | 认证绕过 |
| 3 | GET | `/api3/userinfo` | 过度数据暴露 |
| 4 | GET | `/api4/nolimit` | 资源不限制 |
| 5 | GET | `/api5/admin` | 函数级授权绕过 |
| 6 | POST | `/api6/profile` | 批量赋值漏洞 |
| 7 | GET | `/api7/debug` | 安全配置不当 |
| 8 | POST | `/api8/search` | 注入漏洞 |
| 9 | GET | `/api9/old-api` | 资产管理不当 |
| 10 | POST | `/api10/transfer` | 日志和监控不足 |

## 示例请求

```bash
# 获取物品信息
curl http://127.0.0.1:5019/api1/items/123

# 登录
curl -X POST http://127.0.0.1:5019/api2/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"123456"}'

# 获取用户信息
curl http://127.0.0.1:5019/api3/userinfo
```

## 依赖

- `github.com/gin-gonic/gin` - Web 框架

安装依赖：
```bash
go mod download
```

## 许可证

此项目用于教育和安全研究目的。
