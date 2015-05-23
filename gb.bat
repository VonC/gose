@echo off
setlocal
set GOPATH=%~dp0/deps
set GOBIN=%~dp0/bin
set PATH=%GOROOT%\bin;%PATH%
cd ascom
go version
go build
if %errorlevel% == 0 (
	echo ok
)
endlocal
