# go-specs-greet

学习[验收测试](https://quii.gitbook.io/learn-go-with-tests/testing-fundamentals/scaling-acceptance-tests)。

这章的内容其实就是DDD架构的一种实现（六边形架构）。

```
.
├── adapter_learning_test
│   └── go_test.go
├── adapters
│   ├── docker.go                   // 启动docker容器
│   └── httpserver
│       ├── driver.go               // Driver.Greet()向服务器发送请求，获取greeting
│       └── handler.go              // server的handler，这里是读取name?参数并返回internaction.Greet(name)
├── cmd
│   ├── grpcserver
│   │   └── greeter_server_test.go  
│   └── httpserver
│       ├── Dockerfile              // 启动http server
│       ├── greeter_server_test.go  // 用adapters的docker容器，测试http server
│       └── main.go                 // http server
├── domain
│   └── interactions
│       ├── greet.go                // 核心逻辑(不涉及任何HTTP、Docker、JSON、数据库)，定义了Greet(name string)输出Hello, name
│       └── greet_test.go           // 用specifcation操作规范来测试Greet()
├── go.mod
├── go.sum
├── README.md
└── specifications                  // 操作规范。独立的包，不依赖任何人（定义接口，不引用实现，依赖倒置）
    ├── adapters.go                 // 适配器模式，实现了Greeter接口
    └── greet.go                    // 声明了Greeter接口，并定义了接口Greeter的Greet方法的测试方法是调用greeter.Greet(name)并期待得到Hello, Name
```

