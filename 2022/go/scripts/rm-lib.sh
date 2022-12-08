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

# check directory exists
lib=$(printf "day-%02d\n" "$1")
if [ ! -d "$lib" ]; then
    echo "error: Directory does not exists for $lib" >&2
    exit 1
fi

# remove directory
go work edit -dropuse="$lib"
rm -rf "$lib"
