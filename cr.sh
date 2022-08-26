#!/bin/bash

cat /dev/stdin > $1
go run $1
git add $1
