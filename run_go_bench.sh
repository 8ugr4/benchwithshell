#!/bin/sh
LOG_FILE="test.out"
go test -bench=. > $LOG_FILE
go run ./internal/parse.go
# chmod 755 <script-name>.sh -> makes the text file executable
