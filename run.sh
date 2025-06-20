#!/bin/bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=myuser
export DB_PASS=mypassword
export DB_NAME=marketplace
export SERVER_PORT=8080
export JWT_SECRET=pelindo888

go run cmd/main/main.go 