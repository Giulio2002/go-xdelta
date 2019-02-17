@echo off

setlocal

call "C:\Program Files (x86)\Microsoft Visual Studio 14.0\VC\vcvarsall.bat" amd64

cl.exe /nologo /I src /MT /LD /GL /Fe:go-xdelta-lib.dll xdelta.cpp xdelta-encoder.cpp xdelta-go.cpp /link /RELEASE /LTCG /NOLOGO /VERSION:3.1

del /q go-xdelta-lib.lib go-xdelta-lib.exp *.obj

REM dumpbin.exe /nologo /exports go-xdelta-lib.dll

powershell.exe -ExecutionPolicy Unrestricted -File build-windows-write-version.ps1

endlocal
