#!/bin/bash
set -e
GOOS=linux GOARCH=arm GOARM=7 go build
scp gbis-frame $1:~/
#scp gbis-frame.service $1:~/
