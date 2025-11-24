# API Security OWASP Top 10 2023 POC

本目录用于API安全产品能力检测，覆盖OWASP 2023 API Security Top 10。

## 目录结构
- server/  服务端代码（Flask实现，10个API端点）
- client/  客户端代码（Python脚本，发起Top10攻击/异常请求）

## 快速开始

### 1. 安装依赖
```bash
cd server
pip install -r requirements.txt
```

### 2. 启动服务端
```bash
python app.py
```

### 3. 运行客户端测试
```bash
cd ../client
pip install -r requirements.txt
python run_all.py
```

## 扩展说明
- 可根据实际需求扩展API端点和攻击脚本。
- 详细用法见各目录下README。

## 参考
- [OWASP API Security Top 10 2023](https://owasp.org/API-Security/editions/2023/en/0x08-t10/)
