test: generate
	go test -v -race ./...
cover: generate
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
generate:
	go generate