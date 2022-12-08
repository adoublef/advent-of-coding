#!/bin/bash

main() {
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

    func_asset "$lib" && func_lib "$lib"
}

func_lib() {
    # remove directory
    go work edit -dropuse="$1"
    rm -rf "$1"
}

func_asset() {
    rm "./assets/$1.txt"
}

main "$@"
