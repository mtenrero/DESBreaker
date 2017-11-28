#!/bin/bash

env GOOS=windows GOARCH=amd64 go build -o releases/desbreaker-win-amd64.exe des.go main.go dic.go
env GOOS=windows GOARCH=386 go build -o releases/desbreaker-win-i386.exe des.go main.go dic.go

env GOOS=linux GOARCH=amd64 go build -o releases/desbreaker-linux-i386 des.go main.go dic.go
env GOOS=linux GOARCH=386 go build -o releases/desbreaker-linux-i386 des.go main.go dic.go

env GOOS=darwin GOARCH=amd64 go build -o releases/desbreaker-osx-amd64 des.go main.go dic.go