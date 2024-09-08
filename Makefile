# Change these variables as necessary.
main_package_path = ./cmd/skripsi-be
binary_name = skripsi-be

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^## //p' $(MAKEFILE_LIST) | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

.PHONY: no-dirty
no-dirty:
	@git status --porcelain

# ==================================================================================== #
# Development
# ==================================================================================== #
## install-tools: install development tools for run locally
.PHONY: install-tools
install-tools:
	@echo 'Installing CompileDaemon...'
	@go install github.com/githubnemo/CompileDaemon@v1.4.0

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
