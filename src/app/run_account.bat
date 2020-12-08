@echo off
set ver=1
set server=account.exe
set appName=account
set logflag=1
set pprofIp=127.0.0.1
set addr=127.0.0.1 

set GOPATH=%~dp0
if exist %server% (
    del %server%
)
go build -o %server% main.go
if exist  %server% (
    start %server% -app %appName% -host %addr% -pprof %pprofIp% -log %logflag%
) else (
    echo not find file: %server%
)
