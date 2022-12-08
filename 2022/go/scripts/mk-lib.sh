#!/bin/bash

# check only one argument is allowed
if [[ $# -ne 1 ]]; then
    echo "error: Too many/few arguments, expecting one" >&2
    exit 1
fi

# check argument is a number
re="^[0-9]+$"
if ! [[ $1 =~ $re ]]; then
    echo "error: Not a number" >&2
    exit 1
fi

# check is directory already exists
lib=$(printf "day-%02d\n" "$1")
if [ -d "$lib" ]; then
    echo "error: Directory already exists for $lib" >&2
    exit 1
fi

# create a new library folder
mkdir "$lib"
cd "$lib" || exit 1

go mod init "$lib"
go work use "."

# create two new files inside
echo "package lib" >"lib.go"
echo "package lib" >"lib_test.go"
