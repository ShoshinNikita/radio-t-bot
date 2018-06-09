@ECHO off
set GOPATH=%cd%
ECHO Errors:
C:%HOMEPATH%/go/bin/megacheck bot dates main dialogs
go vet bot dates main dialogs
ECHO Tests:
go test ./...