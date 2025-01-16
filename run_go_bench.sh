#!/bin/sh
LOG_FILE="test.out"
BENCHFILES_DIR="benchmark"
go test -bench=. "$BENCHFILES_DIR"/* > $LOG_FILE
go run main.go
# chmod 755 <script-name>.sh -> makes the text file executable
