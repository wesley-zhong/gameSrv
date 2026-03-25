set CODE_DIR=..\..\cnfGen
set DATA_DIR=..\..\game-json-data
set LUBAN_DLL=.\Tools\Luban\Luban.dll
set CONF_ROOT=.
dotnet %LUBAN_DLL% ^
    -t server ^
    -c go-json ^
    -d json ^
    --conf %CONF_ROOT%\luban.conf ^
    -x outputCodeDir=%CODE_DIR%\cfg\ ^
    -x outputDataDir=%DATA_DIR%\excel-json ^
    -x lubanGoModule=luban 
pause