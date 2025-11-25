# Flow 场景：语义化转账流程示例

该目录演示一个更贴近真实业务的转账流程，包含以下步骤（语义化路径）。使用三步登录流程和 Bearer Token 认证，展示生产环境的安全认证最佳实践。

## 核心特性

- **三步登录流程**：分离密码提交、SMS 请求、SMS 验证，提高安全性
- **Bearer Token 认证**：使用 `Authorization` 请求头传递 token，而非 URL 查询参数
- **请求体参数**：业务参数在 JSON 请求体中传递，提高 API 安全性
- **多阶段状态机**：转账流程需按顺序经过多个阶段（PIN 验证 → SMS 验证 → 提交）

## API 流程

### 登录流程（三步）
1. `POST /flow/login-step1` - 提交账号密码，返回 `loginSessionId`
2. `POST /flow/login-step2` - 请求短信验证码
3. `POST /flow/login-step3` - 提交短信验证码，获得 `token`

### 业务操作（需在请求头中包含 `Authorization: Bearer <token>`）
1. `GET /flow/userinfo` - 获取用户信息
2. `GET /flow/balance` - 查询余额
3. `POST /flow/query-pin` - 验证查询 PIN（可选）
4. `POST /flow/initiate-transfer` - 发起转账
5. `POST /flow/withdraw-pin` - 验证取款 PIN（阶段 0）
6. `POST /flow/sms-code` - 验证短信码（阶段 1）
7. `POST /flow/submit-transfer` - 最终提交转账（阶段 2）

## 快速开始

### 编译

```bash
# 编译服务端
cd go/flow/server
go mod download
go build -o flow-server

# 编译客户端
cd ../client
go mod download
go build -o flow-client
```

### 运行

```bash
# 启动服务（后台运行）
cd go/flow/server
./flow-server &

# 在另一个终端运行客户端
cd ../client
./flow-client --addr http://127.0.0.1:5060 \
  --username alice --password secret \
  --to bob --amount 100 \
  --query-pin 1234 --withdraw-pin 2345 --sms-code 000000
```

## 使用示例

### 完整流程（包含所有步骤）
```bash
./flow-client --addr http://127.0.0.1:5060 \
  --username alice --password secret \
  --to bob --amount 100 \
  --query-pin 1234 --withdraw-pin 2345 --sms-code 000000
```

### 跳过可选步骤（不验证 query-pin）
```bash
./flow-client --addr http://127.0.0.1:5060 \
  --username alice --password secret \
  --to bob --amount 50 \
  --query-pin ""
```

### 测试错误场景

错误的 withdraw PIN：
```bash
./flow-client --addr http://127.0.0.1:5060 \
  --username alice --password secret \
  --to bob --amount 50 \
  --withdraw-pin 9999
```

错误的 SMS 码：
```bash
./flow-client --addr http://127.0.0.1:5060 \
  --username alice --password secret \
  --to bob --amount 50 \
  --sms-code 111111
```

## 可配置参数

| 参数 | 默认值 | 说明 |
|------|--------|------|
| `--addr` | `http://127.0.0.1:5060` | 服务器地址 |
| `--username` | `alice` | 登录用户名 |
| `--password` | `secret` | 登录密码 |
| `--to` | `bob` | 转账目标用户 |
| `--amount` | `100.0` | 转账金额 |
| `--query-pin` | `1234` | 查询 PIN（留空则跳过） |
| `--withdraw-pin` | `2345` | 取款 PIN |
| `--sms-code` | `000000` | 短信验证码 |

## 实现说明

- **认证方式**：三步登录避免在单个请求中暴露密码，符合生产环境最佳实践
- **Token 传递**：使用标准 `Authorization: Bearer <token>` 请求头，而非 URL 查询参数
- **参数传递**：敏感业务参数在请求体中，而非 URL 中，提高安全性
- **状态管理**：内存中维护登录会话和转账状态，便于演示和调试
- **验证码/PIN**：使用固定值（000000、1234、2345），仅用于演示目的

可用命令示例（复制粘贴运行）
```bash
# 正常流程
./flow-client --addr http://127.0.0.1:5060 --username alice --password secret --to bob --amount 50

# 模拟错误的 withdraw PIN
./flow-client --addr http://127.0.0.1:5060 --username alice --password secret --to bob --amount 50 --withdraw-pin 9999

# 模拟错误的 SMS code
./flow-client --addr http://127.0.0.1:5060 --username alice --password secret --to bob --amount 50 --sms-code 111111

# 跳过 query-pin 步骤
./flow-client --addr http://127.0.0.1:5060 --username alice --password secret --to bob --amount 50 --query-pin ""

# 所有参数自定义
./flow-client --addr http://127.0.0.1:5060 --username alice --password secret --to bob --amount 100 --query-pin 1234 --withdraw-pin 2345 --sms-code 000000
```