@echo off

set natsHost=0.0.0.0
set natsPort=23800

nats-server.exe -a %natsHost% -p %natsPort%