# Java服务端（Spring Boot）

## 运行方式

1. 安装JDK 1.8
2. 进入server目录，编译并运行：

```bash
mvn clean package
java -jar target/api-owasp-server-1.0-SNAPSHOT.jar
```

默认监听8080端口。

## API端点
- /api1/items/{itemId}
- /api2/login
- /api3/userinfo
- /api4/nolimit
- /api5/admin
- /api6/profile
- /api7/debug
- /api8/search
- /api9/old-api
- /api10/transfer

每个端点对应OWASP Top 10风险点，详见控制器代码。
