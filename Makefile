PROTO := $(shell find . -type f -name '*.proto')

.phony = protobuf reader writer

all: reader writer

clean:
	rm -f reader writer *.pb.bin

writer: protobuf cmd/writer/*.go
	go build -o writer ./cmd/writer

reader: protobuf cmd/reader/*.go
	go build -o reader ./cmd/reader

protobuf:
	protoc --go_out=. pb/*.proto
	protoc --go_out=. pb/v1/*.proto
	protoc --go_out=. pb/v2/*.proto

test:
	./writer
	./reader *.pb.bin
