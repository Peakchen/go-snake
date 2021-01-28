set DIR=%~dp0
cd /d "%DIR%"
cd proto

echo pb make start...
setlocal enabledelayedexpansion

protoc -I %DIR%proto\ --go_out=plugins=grpc:%DIR%go\ rpc.proto

cd ..
cd .\go
xcopy msgProto %DIR%..\msgProto\ /D /I /F /Y /S/E
rd /s/q msgProto

cd ..\..\..\

pause