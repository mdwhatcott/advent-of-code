#!/usr/bin/env bash

set -x -e

find . -type f -name "go.mod" -execdir ../../gotest.sh \;
