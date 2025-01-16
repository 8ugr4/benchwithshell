#!/bin/sh
LOG_FILE="test.out"
go test ./... -bench=. -benchmem | grep -E '^Bench' > $LOG_FILE
go run main.go
# chmod 755 <script-name>.sh -> makes the text file executable
