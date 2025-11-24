# API Security OWASP Top 10 POC (Java版)

本目录为Java实现，适用于仅有JDK 1.8环境的服务器。包含2019版和2023版的服务端与客户端，覆盖OWASP API Security Top 10。

## 目录结构

```
java/
├── 2019/           # OWASP 2019 API Top 10 Java实现
│   ├── server/     # 服务端（Spring Boot，10个API端点）
│   └── client/     # 客户端（Java，发起Top10攻击/异常请求）
├── 2023/           # OWASP 2023 API Top 10 Java实现
│   ├── server/     # 服务端（Spring Boot，10个API端点）
│   └── client/     # 客户端（Java，发起Top10攻击/异常请求）
└── README.md       # 本说明文件
```

## 快速开始

以2023年POC为例，2019年用法类似：

### 1. 启动服务端
```bash
cd 2023/server
mvn clean package
java -jar target/api-owasp-server-2023-1.0-SNAPSHOT.jar
```

### 2. 运行客户端
```bash
cd ../client
javac OwaspApi2023Client.java
java OwaspApi2023Client
```

## 主要功能
- 覆盖OWASP API Security Top 10（2019/2023）全部风险点
- 每个风险点有独立API端点和对应攻击/异常请求脚本
- 便于API安全产品检测、演示和能力验证

## 扩展说明
- 可根据实际需求扩展API端点和攻击脚本
- 详细用法见各子目录下README

## 参考
- [OWASP API Security Top 10 2019](https://owasp.org/www-project-api-security/)
- [OWASP API Security Top 10 2023](https://owasp.org/API-Security/editions/2023/en/0x08-t10/)

---

如有问题或建议，欢迎提issue或PR。
