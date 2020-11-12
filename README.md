# Simple Protobuf `Any` testing


## Building

Environment:

* `go` `1.15.2`
* `protoc` `1.32`
* `protoc-gen-go` `1.3.1`


```
$ make
protoc --go_out=. pb/*.proto
protoc --go_out=. pb/v1/*.proto
protoc --go_out=. pb/v2/*.proto
go build -o test/writer ./cmd/writer
go build -o test/reader ./cmd/reader
```

## Test result

```
$ make test
[...]
cd test && ./writer && ./reader *.pb.bin
Wrote v2_foo
Wrote no_details
Wrote v1_bar
Wrote v1_foo
--------------------------------------------------------------------------------
no_details.pb.bin
--------------------------------------------------------------------------------
(pb.Payload)value:42

--------------------------------------------------------------------------------
v1_bar.pb.bin
--------------------------------------------------------------------------------
(pb.Payload)value:1 details:<type_url:"type.googleapis.com/pb.v1.Bar" value:"\n\021the bar is closed" >
(*v1.Bar)(0xc000091320)(error:"the bar is closed" )

--------------------------------------------------------------------------------
v1_foo.pb.bin
--------------------------------------------------------------------------------
(pb.Payload)value:1 details:<type_url:"type.googleapis.com/pb.v1.Foo" value:"\n\016this is foo v1" >
(*v1.Foo)(0xc0000916b0)(message:"this is foo v1" )
(*v2.Foo)(0xc0000e2880)(message:"this is foo v1" )

--------------------------------------------------------------------------------
v2_foo.pb.bin
--------------------------------------------------------------------------------
(pb.Payload)value:2 details:<type_url:"type.googleapis.com/pb.v2.Foo" value:"\n\016this is foo v2\021\000\000\000\000\000\000\360?" >
(*v1.Foo)(0xc000091c80)(message:"this is foo v2" 2:4607182418800017408 )
(*v2.Foo)(0xc0000e2a80)(message:"this is foo v2" ratio:1 )
```
