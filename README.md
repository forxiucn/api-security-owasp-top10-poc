# API Security OWASP Top 10 POC

本项目用于API安全产品能力检测，覆盖OWASP API Security Top 10（2019 & 2023），支持多语言实现（Python、Java、Golang）。通过模拟常见API安全风险点，帮助安全产品进行能力验证和演示。

## 目录结构
```
api-security-owasp-top10/
├── 2019/           # Python实现（2019版）
├── 2023/           # Python实现（2023版）
├── java/           # Java实现（2019/2023）
├── go/             # Golang实现（2019/2023）
└── README.md       # 项目说明
```

## 快速开始

以Golang 2023年POC为例，其他语言和版本用法类似：

### 1. 启动服务端
```bash
cd go/2023/server
# 安装依赖
go mod tidy
# 启动服务
go run main.go
# 或编译为可执行文件
go build -o server_linux main.go
./server_linux
```

### 2. 运行客户端
```bash
cd ../client
# 运行测试
go run main.go
```

## 主要功能
- 覆盖OWASP API Security Top 10（2019/2023）全部风险点
- 每个风险点有独立API端点和对应攻击/异常请求脚本
- 支持Python、Java、Golang多种实现
- 便于API安全产品检测、演示和能力验证

## 扩展说明
- 可根据实际需求扩展API端点和攻击脚本
- 详细用法见各子目录下README

## 参考
- [OWASP API Security Top 10 2019](https://owasp.org/www-project-api-security/)
- [OWASP API Security Top 10 2023](https://owasp.org/API-Security/editions/2023/en/0x08-t10/)

---
如有问题或建议，欢迎提issue或PR。
