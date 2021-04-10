REM  This script assumes the file icon.ico exists and is up to date.
REM  It does not use icon.png. To convert that to icon.ico the current process is to use a service like http://icoconvert.com to make a multi-resolution icon file from the png.
REM  This script creates the .syso files that the Go build tool uses to give the resulting .exe its icon.

go get github.com/gonutz/rsrc
rsrc -arch=386   -ico=icon.ico -o=ide_386.syso
rsrc -arch=amd64 -ico=icon.ico -o=ide_amd64.syso