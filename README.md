# Catalog Microservices Sample Application

对当下自己项目中的业务，进行一个微服务改造，需要考虑如下技术点：

- 微服务架构（BFF、Service、Admin、Job、Task 分模块）
- API 设计（包括 API 定义、错误码规范、Error 的使用）
- gRPC 的使用
- Go 项目工程化（项目结构、DI、代码分层、ORM 框架）
- 并发的使用（errgroup 的并行链路请求）
- 微服务中间件的使用（ELK、Opentracing、Prometheus、Kafka）
- 缓存的使用优化（一致性处理、Pipeline 优化）

## Architecture Diagram

![alt text](catalog_diagram.png 'Catalog Diagram')

### Questions

请问要如何在 Kratos 框架中引入消息队列？
我在 catalog-job 尝试开了一个 job/task 的 grpc 接口来开启 consumer job，不过我认为这不是一个好方案。
我的想法法案有几个：
1） 直接使用 kafka client 来做数据库应用 2) 开启 kratos job service 时，start consumer in background。
