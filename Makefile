format: ## Formats the code. Must have goimports installed (use make install-linters or make install-imports).
	goimports -w -local github.com/Nikit123/snake_game ./pkg

install-linters: ## Installs go-linter
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

install-imports: ## Installs goimports
	go get golang.org/x/tools/cmd/goimports