build:
	@go build -o untgz

run:
	@./untgz

test:
	@go test ./...