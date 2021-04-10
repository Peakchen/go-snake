
@echo off
set DIR=%~dp0
cd /d "%DIR%"
cd pb
cd ..
del *.pb
echo clean client/*.pb

cd proto

setlocal enabledelayedexpansion
for /r %%i in (*.proto) do ( 
      set pbname=%%i 
      set pbname=!pbname:~0,-6!pb
      protoc -I %DIR% --descriptor_set_out !pbname! %%i 
	  protoc --version
)

xcopy *pb client /D /I /F /Y
xcopy *pb ..\pb /D /I /F /Y

dir client
del *.pb
del /F /S /Q client
rd /S /Q client 

echo "Done"
PAUSE