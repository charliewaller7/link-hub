test:
	@go test -test.v ./...

build:
	@go build -o linkhub .

cross-comp:
	@echo "cross-compile"