#!/bin/bash
set -e
GOOS=linux GOARCH=arm go build
scp gbis-frame $1:~/
#scp -r _python $1:~/