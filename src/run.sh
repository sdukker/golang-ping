#!/bin/sh

go mod download
go build -o /usr/src/build/build
/usr/src/build/build