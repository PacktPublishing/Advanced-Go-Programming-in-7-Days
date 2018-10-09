.PHONY: format
format:
	@find . -type f -name "*.go*" -print0 | xargs -0 gofmt -s -w

.PHONY: debs
debs:
	@go get -u github.com/golang/dep/cmd/dep
	@dep ensure

.PHONY: test
test:
	@go test -race ./...

.PHONY: bench
bench:
	@go test -bench=. -benchmem

# Clean junk
.PHONY: clean
clean:
	@go clean ./...