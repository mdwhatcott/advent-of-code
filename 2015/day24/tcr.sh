#!/usr/bin/env bash

function test() {
    date > out
    go test -v >> out 2>&1
}
function commit() {
    git add -A
    git commit -m "tcr"
}
function revert() {
    git reset --hard
}

test && commit || revert
