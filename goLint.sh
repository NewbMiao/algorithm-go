#!/usr/bin/env bash
workspace=$(cd $(dirname $0) && pwd -P)
# eg: 
# ./goLint.sh # lint all files
# ./goLint.sh leetcode/... # lint files in leetcode
# ./goLint.sh leetcode/... --fix # lint files in leetcode and autofix
docker run --rm -v "$workspace":/app -w /app golangci/golangci-lint golangci-lint run $1 $2
