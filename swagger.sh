#!/bin/bash

go install github.com/swaggo/swag/cmd/swag@latest
$GOBIN/swag init --dir ./cmd,./
