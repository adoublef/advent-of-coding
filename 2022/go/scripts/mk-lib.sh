#!/bin/bash

func_main() {
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

    day=$1

    # check is directory already exists
    lib=$(printf "day-%02d\n" "$1")
    if [ -d "$lib" ]; then
        echo "error: Directory already exists for $lib" >&2
        exit 1
    fi

    func_asset "$day" "$lib" && func_lib "$lib"
}

func_lib() {
    # create a new library folder
    mkdir "$1"
    cd "$1" || exit 1

    go mod init "$1"
    go work use "."

    # create two new files inside
    echo "package lib" >"lib.go"
    echo "package lib" >"lib_test.go"

    code lib{,_test}.go
}

func_asset() {
    content=$(curl "https://adventofcode.com/2022/day/$1/input" \
        -H 'authority: adventofcode.com' \
        -H 'accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9' \
        -H 'accept-language: en-GB,en;q=0.9,en-US;q=0.8' \
        -H 'cache-control: max-age=0' \
        -H 'cookie: session=53616c7465645f5f285508215ac3158dd88411e7fd96d7d28bc836435794e6d0032aaf2276f2b53152b66a577c206f2b36242addb0f6277ea1667b3c4e069383' \
        -H 'dnt: 1' \
        -H 'referer: https://adventofcode.com/2022/day/6' \
        -H 'sec-ch-ua: "Not?A_Brand";v="8", "Chromium";v="108", "Microsoft Edge";v="108"' \
        -H 'sec-ch-ua-mobile: ?0' \
        -H 'sec-ch-ua-platform: "Windows"' \
        -H 'sec-fetch-dest: document' \
        -H 'sec-fetch-mode: navigate' \
        -H 'sec-fetch-site: same-origin' \
        -H 'sec-fetch-user: ?1' \
        -H 'upgrade-insecure-requests: 1' \
        -H 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 Edg/108.0.1462.42' \
        --compressed --silent --show-error)

    not_found="404 Not Found"
    unavaible="Please don't repeatedly request this endpoint before it unlocks!"

    if [[ "$content" == *"$not_found"* || "$content" == *"$unavaible"* ]]; then
        echo "error: Not found" >&2
        exit 1
    fi

    touch "assets/$2.txt"
    echo "$content" >"assets/$2.txt"
}

func_main "$@"
