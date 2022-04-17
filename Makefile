GO = go

.PHONY: deps
deps: go.mod

go.mod:
	$(GO) mod init
	$(GO) mod tidy

.PHONY: clean
clean:
	$(GO) clean
