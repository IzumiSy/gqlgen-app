#!/bin/bash

go generate ./...
go run ./server.go
