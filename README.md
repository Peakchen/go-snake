# go-snake

![](https://img.shields.io/github/issues/Peakchen/go-snake)	![](https://img.shields.io/github/forks/Peakchen/go-snake)![](https://img.shields.io/github/stars/Peakchen/go-snake)![](https://img.shields.io/github/license/Peakchen/go-snake)

#### 软件架构

1. gate server （内、外网管）数据安全防护，负载均衡, 分发客户端消息
2. game server 消息处理，数据存储 
3. login server 登录服(账号登录，微信登录...)
4. robot server 客户端模拟器
5. db mysql 
6. 网络协议采用protobuf（后端：go，前端：js）
7. simulation 服务器模拟器
8. etcd 服务发现
9. grpc 服务期间远程调用
10. 后台控制服务器（CURD后台数据，采用iris 开源web 框架）
11. 数据存储、逻辑处理模块等独立，可按需分配到到各服使用
12. nats 消息队列基于网关与逻辑服间通信
13. world 世界服aoi（实现中...）

#### 使用说明
serverApp 参数列表：   
"args": [
    "-app",
    "gate",
    "-ver",
    "1",
    "-webhost",
    "127.0.0.1:19921",
    "-tcphost",
    "127.0.0.1:19922",
    "-pprof",
    "127.0.0.1",
    "-log",
    "1",
    "-clients",
    "1"
]

统一从main入口，各个服务器模块各有init，run等初始启动方式。

<img src="https://github.com/Peakchen/go-snake/tree/master/src/app/note/struct.PNG" width="50%">