set DIR=%~dp0
cd /d "%DIR%"
cd proto

echo pb make start...
setlocal enabledelayedexpansion
for %%i in (*.proto) do ( 
	  echo %%i 
	  set jsname=%DIR%js\ 
	  protoc --js_out=import_style=commonjs,binary:!jsname! %%i
)

pause
echo "Done"
PAUSE