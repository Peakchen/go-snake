#!/bin/bash

protoc -I=./ --js_out=import_style=commonjs,binary:./client_js ./*.proto

cd client_js

browserify exportPb.js -o  exportPb_web.js

#for file in `ls *.js`
#do
    #echo $file
    #watchify $file -o  ${file%%.*}_web.js
    #browserify $file -o  ${file%%.*}_web.js
#done