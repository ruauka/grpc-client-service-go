gen:
	@protoc --proto_path=./proto --go_out=. --go-grpc_out=. execute.proto

lint:
	@golangci-lint run
