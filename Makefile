OK_RESULT= echo "Done😌!"
CMD_ROOT_FILE_PATH=./main.go

install:
	@echo "Installing all tools for developing..."
	@sh ./scripts/install.sh && ${OK_RESULT}

tests:
	@echo "Run tests..."
	@go test ./... && ${OK_RESULT}

lint:
	@echo "Run lints..."
	@golangci-lint run
	@gosec ./...

benchs:
	@echo "Run benchmark..."
	@go clean -testcache && go test -run=XXX -bench=.
