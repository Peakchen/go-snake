##Software architecture
1. Data security protection of gate server, load balancing and distribution of client messages
2. Game server message processing, data storage
3. Login server login service (for common login [by accountID and pwd] and wechat)
4. Robot server client simulator
5. db mysql
6. The network protocol adopts protobuf (back end: go, front end: JS)
7. simulation server simulator
8. etcd service discovery
9. Remote call during grpc service
10. The nats message queue is based on the communication between the gateway and the logical server (planned) 

##Instructions for use

Serverapp parameter list:  

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


All server modules have init, run and other initial start modes.

<img src="https://github.com/Peakchen/go-snake/src/app/note/struct.png" width="50%">