REM  This script turns the .ttf font into Go source code so that the data is part of the built code. This means that the font file does not have to be loaded at runtime.
REM  The Go fonts are taken from https://github.com/golang/image/blob/master/font/gofont/ttfs

go get github.com/gonutz/bin2go/v2/bin2go
bin2go -package=font -var=TTF < Go-Mono.ttf > font.go