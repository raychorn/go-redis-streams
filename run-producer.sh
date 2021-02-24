#!/bin/bash

export PATH=$GOPATH/bin:$PATH

cd /workspaces
go run producer/main.go
