# todolist-ddd
基于 DDD(Domain Driven Design) 领域驱动设计 架构实现todolist. 

在此之前我查阅了很多资料，每个人都有自己理解的DDD。

这里我也只说我自己理解的 DDD，如何和你所理解的有出入，那一定是你对，我理解的是错的。

# 架构
```shell
./todolist-ddd
├── application         // 应用层: 做domain编排
│   ├── task            // task 应用层模块
│   └── user            // user 应用层模块
├── cmd                 // 启动入口
├── conf                // 配置文件
├── consts              // 常量定义
├── docs                // 接口文档
├── domain              // 领域层: 
│   ├── task            // task 领域层模块
│   │   ├── entity      // task 实体定义及充血对象
│   │   ├── repository  // task 实体的数据持久化接口
│   │   └── service     // task 具体业务逻辑
│   └── user            // user 领域层模块
│      ├── entity       // user 实体定义及充血对象
│      ├── repository   // user 实体的数据持久化接口
│      └── service      // user 具体业务逻辑
├── infrastructure      // 基础架构层: 提供数据来源和基础服务能力
│   ├── auth            // 鉴权认证服务
│   ├── common          // 公共服务
│   │   ├── context     // context 上下游管理
│   │   └── log         // log 服务
│   ├── encrypt         // 加密 服务
│   └── persistence     // 持久层
│       ├── dbs         // db数据连接
│       ├── task        // task 的dao层 访问task数据库
│       └── user        // user 的dao层 访问user数据库
├── interfaces          // 接口层: 对接不同的端进行适配转化
│   ├── adapter         // 适配器
│   │   └── initialize  // Web 路由初始化
│   ├── controller      // controller 层
│   ├── midddleware     // 中间件
│   └── types           // 类型
└── logs                // 日志文件存储
```
