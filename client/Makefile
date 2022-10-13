test:
	go test -cover ./...

lint:
	@golangci-lint run

vendor:
	go mod vendor

gen:
	@protoc --proto_path=./api/proto --go_out=. --go-grpc_out=. execute.proto