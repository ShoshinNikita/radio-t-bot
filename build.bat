@ECHO off
set GOPATH=%cd%
CALL test

SET /p command=Should continue(y/n): 
IF "%command%" == "y" (
	SET GOOS=linux
	SET GOARCH=amd64
	go build -o app/radio-t-bot src/main/main.go

	docker build -t radio-t-bot .
)