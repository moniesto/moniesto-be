#!/bin/bash

go install github.com/swaggo/swag/cmd/swag@latest
swag init --dir ./cmd,./
