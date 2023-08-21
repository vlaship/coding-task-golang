tests:
	@echo "Running tests"
	@go test -v ./...

upgrade:
	@echo "Upgrading dependencies"
	@go get -u ./...
	@go mod tidy
