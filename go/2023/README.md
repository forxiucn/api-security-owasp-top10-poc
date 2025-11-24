# OWASP API Security Top 10 2023 (Golang)

本目录为2023年OWASP API Top 10的Golang实现，包含服务端和客户端。

## 目录结构
```

## 编译为可执行文件

可根据目标平台打包二进制文件，例如：

### Linux 64位
```bash
GOOS=linux GOARCH=amd64 go build -o server_linux main.go
```
### Windows 64位
```bash
GOOS=windows GOARCH=amd64 go build -o server_windows.exe main.go
```
### Mac 64位
```bash
GOOS=darwin GOARCH=amd64 go build -o server_mac main.go
```
客户端同理，将 main.go 替换为客户端入口文件。
2023/
├── server/   # Gin服务端，10个API端点，监听5023端口
├── client/   # Go客户端，依次请求10个端点
```

## 快速开始

### 1. 启动服务端
```bash
cd server
# 安装依赖
go mod tidy
# 启动服务
go run main.go
```

### 2. 运行客户端
```bash
cd ../client
# 运行测试
go run main.go
```

## 功能说明
- 服务端实现OWASP 2023 Top 10全部风险点API
- 客户端依次发起攻击/异常请求，便于安全产品检测

---
如需扩展API或脚本，直接修改main.go即可。
