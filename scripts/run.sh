#!/bin/bash

set -a; eval "$(cat .devcontainer/.env <(echo) <(declare -x))"; set +a
go run main.go
