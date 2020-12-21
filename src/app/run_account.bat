@echo off
set ver=1
set server=account.exe
set appName=account
set logflag=1
set pprofIp=127.0.0.1:18000
set webhost=127.0.0.1:19921
set tcphost=127.0.0.1:19922

set GOPATH=%~dp0
if exist %server% (
    del %server%
)
go build -o %server% main.go
if exist  %server% (
    start %server% -app %appName% -ver %ver% -webhost %webhost% -tcphost %tcphost% -pprof %pprofIp% -log %logflag%
) else (
    echo not find file: %server%
)
