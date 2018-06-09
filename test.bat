@ECHO off
set GOPATH=%cd%
ECHO Errors:
C:%HOMEPATH%/go/bin/megacheck bot dates main
go vet bot dates main
ECHO Tests:
go test ./...