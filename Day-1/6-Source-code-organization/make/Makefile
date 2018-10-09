APP=app

all: fmt test build

.PHONY: run
run: build
	./$(APP)

.PHONY: doc
doc:
	godoc -http=:6060 -index

.PHONY: build
build:
	go build -o $(APP) -v ./...

.PHONY: test
test:
	go test -timeout 20s -race -v ./...

.PHONY: debs
debs:
	go get ./...

.PHONY: fmt
fmt:
	gofmt -l -s .

.PHONY: clean
clean:
	go clean
	rm -f $(APP)