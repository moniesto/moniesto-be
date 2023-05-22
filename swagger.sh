#!/bin/bash

go install github.com/swaggo/swag/cmd/swag@v1.8.12
swag init --dir ./cmd,./
