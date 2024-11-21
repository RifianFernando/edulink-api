# applicationName.
main_package_path = ./cmd/edulink-api
binary_name = edulink-api

# ==================================================================================== #
# Development
# ==================================================================================== #
## install-tools: install development tools for run locally
.PHONY: install-tools
install-tools:
	@echo 'Installing Go mod...'
	@go mod tidy

	@echo 'Installing CompileDaemon...'
	@go install github.com/githubnemo/CompileDaemon@v1.4.0

	@echo 'Copy .env.example to .env...'
	@copy .env.example .env

	@echo 'Finished installing development tools.'

## build: build the application
.PHONY: build
build:
	# Include additional build steps, like TypeScript, SCSS or Tailwind compilation here...
	@go build -o=/tmp/bin/${binary_name} ${main_package_path}

# ==================================================================================== #
# Deployment
# ==================================================================================== #
## production/deploy: deploy the application to production
.PHONY: production/deploy
production/deploy: confirm no-dirty
	GOOS=linux GOARCH=amd64 go build -ldflags='-s' -o=/tmp/bin/linux_amd64/${binary_name} ${main_package_path}
	# upx -5 /tmp/bin/linux_amd64/${binary_name}
	# Include additional deployment steps here...
