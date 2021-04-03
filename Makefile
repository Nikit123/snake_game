format: ## Formats the code. Must have goimports installed (use make install-linters or make install-imports).
	goimports -w -local github.com/Nikit123/snake_game ./pkg

install-linters:
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

install-imports:
	go get golang.org/x/tools/cmd/goimports