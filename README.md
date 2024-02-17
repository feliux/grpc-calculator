# gRPC Calculator

Simple gRPC calculator implementation with SSL, login and headers interceptors.

**gRPC Unary**: 
- `rpc Sum(SumRequest) returns (SumResponse);`
- Error Handling: `rpc Sqrt(SqrtRequest) returns (SqrtResponse);`
**gRPC Server Streaming**: `rpc Primes(PrimesRequest) returns (stream PrimesResponse)`
**gRPC Client Streaming**: `rpc Average(stream AverageRequest) returns (AverageResponse)`
**gRPC Bidirectional**: `rpc Max(stream MaxRequest) returns (stream MaxResponse)`

## Usage

1. Install protobuf go plugins

```bash
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
$ export PATH="$PATH:$(go env GOPATH)/bin"
```

2. Compile proto files

```bash
$ cd proto
$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative *.proto
```

3. Uncomment [client/main.go](./client/main.go) with the function to test

```
//DoSum(c, &pb.SumRequest{FirstValue: 1, SecondValue: 1})
//DoSumDeadline(c, &pb.SumRequest{FirstValue: 1, SecondValue: 1}, 5*time.Second) // no error
//DoSumDeadline(c, &pb.SumRequest{FirstValue: 1, SecondValue: 1}, 1*time.Second) // timeout error
//DoPrimes(c, 400)
//DoAverage(c, []float64{1, 2, 3, 4})
//DoMax(c, []int32{1, 5, 3, 6, 2, 20})
//DoSqrt(c, -9)
```

4. Generate certs

```bash
$ cd ssl
$ bash generate-certs.sh
```

5. Compile server and client

```bash
$ cd ..
$ go build -o server/server server/*.go
$ go build -o client/client client/*.go
```

6. Run server and client

```bash
# Run server
$ ./server/server
# Run client on another terminal
$ ./client/client
```
