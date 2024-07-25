OK_RESULT= echo "Done😌!"
CMD_ROOT_FILE_PATH=./main.go

install:
	@echo "Installing all tools for developing..."
	@sh ./scripts/install.sh && ${OK_RESULT}

build: ## Build application
	@echo "Building application..."
	@go build -o dist/godom $(CMD_ROOT_FILE_PATH) && ${OK_RESULT}
