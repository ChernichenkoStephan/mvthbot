#!/usr/bin/env bash

set -e

# test with -race
clear && go test --timeout 5m -race ./...

