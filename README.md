# go-snake

#### 软件架构

1. gate server 数据安全防护，负载均衡, 分发客户端消息
2. game server 消息处理，数据存储 
3. login server 登录服
4. robot server 客户端模拟器
5. db mysql 
6. ... 

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



