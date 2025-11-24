# Java服务端（Spring Boot，2023版）

## 运行方式

1. 安装JDK 1.8
2. 进入server目录，编译并运行：

```bash
mvn clean package
java -jar target/api-owasp-server-2023-1.0-SNAPSHOT.jar
```

默认监听8080端口。

## API端点
- /api1/items/{itemId}  # BOLA
- /api2/login           # Broken Authentication
- /api3/userinfo        # BOPLA
- /api4/nolimit         # Unrestricted Resource Consumption
- /api5/admin           # Broken Function Level Authorization
- /api6/transfer        # Unrestricted Access to Sensitive Business Flows
- /api7/ssrf            # SSRF
- /api8/debug           # Security Misconfiguration
- /api9/old-api         # Improper Inventory Management
- /api10/unsafe         # Unsafe Consumption of APIs

每个端点对应2023年OWASP API Top 10风险点，详见控制器代码。
