#!/bin/bash

root=`pwd`

function buildApp(){
    export GOPATH=$root
    GOOS=linux go build -o goSnake main.go 
}

function send(){
    pkg=app-$(date +%Y-%m-%d-%H%M%S).tar.gz
    tar -czf $pkg goSnake ./ini/server.ini
    scp $pkg root@www.peakchen.online:/develop/snake/
    ssh root@www.peakchen.online "cd /develop/snake;tar -zxvf $pkg;rm $pkg -rf;supervisorctl start all"
    rm goSnake -rf
    rm $pkg -rf
}

buildApp
send