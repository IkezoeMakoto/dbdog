DIST=dist
$(DIST):
	mkdir -p $(DIST)

build: $(DIST)
	go build -o $(DIST)/dbdog main.go
.PHONY: build

fmt:
	go fmt ./...
.PHONY: fmt

run:
	go run main.go
.PHONY: run
