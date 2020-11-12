.phony = protobuf

all: build

build: test/writer test/reader

test/writer: protobuf cmd/writer/*.go
	go build -o test/writer ./cmd/writer

test/reader: protobuf cmd/reader/*.go
	go build -o test/reader ./cmd/reader

protobuf:
	protoc --go_out=. pb/*.proto
	protoc --go_out=. pb/v1/*.proto
	protoc --go_out=. pb/v2/*.proto

test: build
	cd test && ./writer && ./reader *.pb.bin

