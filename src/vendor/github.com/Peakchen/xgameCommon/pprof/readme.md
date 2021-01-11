# cpu and memory analysis cmds:

- go-torch -u  http://localhost:12004/debug/pprof/heap --colors mem --raw  -f mem.svg
- go-torch -u  http://localhost:12004 --seconds 60 --raw -f cpu.svg
- go tool pprof -raw -seconds 60 http://localhost:12004/debug/pprof/profile
- go tool pprof http://localhost:12004/debug/pprof/profile

## cpu analysis result:

ExternalGateWayServer

![ExternalGW_cpu](https://github.com/Peakchen/foho/blob/master/src/ExternalGateway/analysis/ExternalGW_cpu_201911291749.png)

InnerGateWayServer

![InnerGW_cpu](https://github.com/Peakchen/foho/blob/master/src/InnerGateway/analysis/InnerGW_cpu_201911291750.png)

GameServer

![gameSvr_cpu](https://github.com/Peakchen/foho/blob/master/src/GameServer/analysis/gameSvr_cpu_201911291746.png)

LoginServer

![loginSvr_cpu](https://github.com/Peakchen/foho/blob/master/src/LoginServer/analysis/loginSvr_cpu_201911291740.png)