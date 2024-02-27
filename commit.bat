@echo off
rem
rem Ejemplo de manejo de la fecha y hora actual - v2014-02-05
rem

chcp 1252 > NUL

setlocal

git add .

SET PARAM=%1
:LOOP
SHIFT
IF %1.==. GOTO JUMP
SET PARAM=%PARAM% %1
GOTO LOOP
:JUMP

set TEXTO=%PARAM% - %DATE% - %TIME%

git commit -m "%TEXTO%"

git push

set GOOS=linux
set GOARCH=amd64
go build -tags lambda.norpc -o bootstrap main.go
del main.zip
tar.exe -a -cf main.zip main

