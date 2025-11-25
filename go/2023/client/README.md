# OWASP API Top 10 2023 Client

用于测试 OWASP API Top 10 2023 漏洞的客户端应用，可攻击并验证演示服务端的安全问题。

## 前置要求

- Go 1.16 或更高版本
- 目标服务端已启动并可访问

## 编译

编译为可执行文件，支持单架构或交叉编译。

### 单架构编译

```bash
go build -o client .
```

### 交叉编译示例

```bash
# macOS Apple Silicon（ARM64）
GOOS=darwin GOARCH=arm64 go build -o 2023_client_mac_arm64 .

# macOS Intel（amd64）
GOOS=darwin GOARCH=amd64 go build -o 2023_client_mac_amd64 .

# Linux
GOOS=linux GOARCH=amd64 go build -o 2023_client_linux_amd64 .

# Windows
GOOS=windows GOARCH=amd64 go build -o 2023_client_amd64.exe .
```

## 运行

编译后运行二进制文件，支持以下参数：

```bash
./client [OPTIONS]
```

### 参数说明

| 参数 | 默认值 | 说明 | 示例 |
|------|--------|------|------|
| `--addr` | `http://127.0.0.1:5023` | 服务器地址，必须包含协议，格式为 `protocol://host:port` 或 `protocol://host:port/contentPath`（例如 `http://127.0.0.1:5023/api`） | `--addr http://api.example.com:5023` |

### 运行示例

```bash
# 默认连接（http://127.0.0.1:5023）
./client

# 指定服务器主机、端口和 content path（包含协议）
./client --addr http://192.168.1.100:5023

./client --addr http://127.0.0.1:8080/v1/api

# 远程服务器 + content path
./client --addr https://api.example.com:443/api/v2023
```

## 测试场景

客户端会自动执行以下 10 个测试用例，对应 OWASP API Top 10 2023 的各个漏洞：

1. **对象级授权绕过 (BOLA)** - 直接访问他人对象
2. **认证绕过** - 测试默认凭证和认证漏洞
3. **对象属性级授权绕过 (BOPLA)** - 获取不应暴露的属性（如薪资）
4. **资源消耗无限制** - 无限制消耗服务器资源
5. **函数级授权绕过** - 访问受限的管理功能
6. **敏感业务流程无限制** - 绕过业务流程控制
7. **服务端请求伪造 (SSRF)** - 让服务器访问任意 URL
8. **安全配置不当** - 访问调试和配置信息
9. **库存管理不当** - 访问已弃用的 API
10. **不安全的 API 使用** - 不安全地调用外部 API

## 示例运行

```bash
# 在本地启动客户端测试
./client

# 输出示例：
# Running client against http://127.0.0.1:5023
# /api1/items/123: 200 {"item_id":"123","detail":"Object info (no auth check)"}
# /api2/login: 200 {"msg":"Login success","token":"fake-jwt-token"}
# /api3/userinfo: 200 {"username":"alice","email":"alice@example.com","role":"admin","salary":10000}
# ...
```

## 与 2019 版本的差异

OWASP API Top 10 2023 相比 2019 版本的客户端测试差异：

- 第 3 项测试：从基本的过度数据暴露升级为对象属性级测试（包括 `salary` 字段暴露）
- 第 6 项测试：从批量赋值升级为敏感业务流程测试（`/api6/transfer`）
- 第 7 项测试：新增 SSRF 漏洞测试（`/api7/ssrf`）
- 第 8 项测试：调整为 `/api8/debug` GET 请求

## 依赖

无外部依赖（使用 Go 标准库）

## 许可证

此项目用于教育和安全研究目的。
