test:
	go test -cover ./...

lint:
	@golangci-lint run

vendor:
	go mod vendor

proto:
	@protoc --proto_path=./api/proto --go_out=. --go-grpc_out=. strategy.proto

bench:
	go test -bench=BenchmarkExecute -benchmem -benchtime 5s -count=5

pprof:
	go test -bench=BenchmarkExecute -benchmem -benchtime=5s -cpuprofile cpu.out -memprofile mem.out

pprof-mem:
	go tool pprof -http :9000 mem.out

pprof-cpu:
	go tool pprof -http :9000 cpu.out