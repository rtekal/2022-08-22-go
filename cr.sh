#!/bin/bash

cat /dev/stdin > $1.go
go run $1.go
git add $1.go
