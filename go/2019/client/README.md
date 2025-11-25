# OWASP API Top 10 2019 Client

用于测试 OWASP API Top 10 2019 漏洞的客户端应用，可攻击并验证演示服务端的安全问题。

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
GOOS=darwin GOARCH=arm64 go build -o 2019_client_mac_arm64 .

# macOS Intel（amd64）
GOOS=darwin GOARCH=amd64 go build -o 2019_client_mac_amd64 .

# Linux
GOOS=linux GOARCH=amd64 go build -o 2019_client_linux_amd64 .

# Windows
GOOS=windows GOARCH=amd64 go build -o 2019_client_amd64.exe .
```

## 运行

编译后运行二进制文件，支持以下参数：

```bash
./client [OPTIONS]
```

### 参数说明

| 参数 | 默认值 | 说明 | 示例 |
|------|--------|------|------|
| `--addr` | `http://127.0.0.1:5019` | 服务器地址，必须包含协议，格式为 `protocol://host:port` 或 `protocol://host:port/contentPath`（例如 `http://127.0.0.1:5019/api`） | `--addr http://api.example.com:5019` |

### 运行示例

```bash
# 默认连接（http://127.0.0.1:5019）
./client

# 指定服务器主机、端口和 content path（包含协议）
./client --addr http://192.168.1.100:5019

./client --addr http://127.0.0.1:8080/v1/api

# 远程服务器 + content path
./client --addr https://api.example.com:443/api/v2019
```

## 测试场景

客户端会自动执行以下 10 个测试用例，对应 OWASP API Top 10 的各个漏洞：

1. **对象级授权绕过 (BOLA)** - 直接访问他人对象
2. **认证绕过** - 测试默认凭证和认证漏洞
3. **过度数据暴露** - 获取不应暴露的敏感字段
4. **资源不限制** - 无速率限制的端点测试
5. **函数级授权绕过** - 访问受限的管理功能
6. **批量赋值漏洞** - 修改非预期的对象属性
7. **安全配置不当** - 访问调试和配置信息
8. **注入漏洞** - SQL 注入和命令注入测试
9. **资产管理不当** - 访问已弃用的 API
10. **日志和监控不足** - 关键业务流程无日志记录

## 示例运行

```bash
# 在本地启动客户端测试
./client

# 输出示例：
# Running client against http://127.0.0.1:5019
# /api1/items/123: 200 {"item_id":"123","detail":"Object info (no auth check)"}
# /api2/login: 200 {"msg":"Login success","token":"fake-jwt-token"}
# ...
```

## 依赖

无外部依赖（使用 Go 标准库）

## 许可证

此项目用于教育和安全研究目的。
