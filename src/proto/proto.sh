#!/bin/bash
#dir=`ls ./*.proto` #定义遍历的目录
#for file in $dir
#do
#    clang-format -i -style="{AlignConsecutiveAssignments: true,AlignConsecutiveDeclarations: true,AllowShortFunctionsOnASingleLine: None,BreakBeforeBraces: GNU,ColumnLimit: 0,IndentWidth: 4,Language: Proto}" $file
#done
#rm -f ../akmessage/*.go
./protoc -I=./ --go_out=plugins=grpc:../akmessage ./*.proto
