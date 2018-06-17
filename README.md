## Running

```
./huestacean-cli
```

For a list of options
```
./huestacean-cli --help
```

## Build

Install Go; minimum tested is 1.10.1Install Go (1.10)

### Dependencies
WIP

### Build

```
go build
```

### Compile protocol buffer files

[Install protoc](https://github.com/google/protobuf/releases) and ensure it's on your PATH

```
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u google.golang.org/grpc
protoc -I ./proto --go_out=plugins=grpc:proto-gen ./proto/*.proto
```