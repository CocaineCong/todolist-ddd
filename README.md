# todolist-ddd
基于 DDD(Domain Driven Design) 领域驱动设计 架构实现todolist. 

在此之前我查阅了很多资料，每个人都有自己理解的DDD。

这里我也只说我自己理解的 DDD，如何和你所理解的有出入，那一定是你对，我理解的是错的。

# 架构
```shell
./todolist-ddd
├── api
├── application
│  ├── task
│  └── user
├── cmd
├── conf
├── domain
│  ├── task
│  │  ├── entity
│  │  ├── repository
│  │  └── service
│  └── user
│     ├── entity
│     ├── repository
│     └── service
├── infrastructure
│  ├── auth
│  ├── common
│  │  ├── context
│  │  └── log
│  ├── consts
│  ├── encrypt
│  └── persistence
│     ├── dbs
│     ├── task
│     └── user
├── interfaces
│  ├── adapter
│  │  └── initialize
│  ├── controller
│  ├── midddleware
│  └── types
└── logs

```