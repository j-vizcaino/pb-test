# Simple Protobuf `Any` testing with lists


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
Wrote v1_foo
Wrote v1_foo_v1_bar
Wrote v1_bar_v2_foo
Wrote no_details
Wrote v1_bar
--------------------------------------------------------------------------------
no_details.pb.bin
--------------------------------------------------------------------------------
(pb.Payload)value:42

--------------------------------------------------------------------------------
v1_bar.pb.bin
--------------------------------------------------------------------------------
(pb.Payload)value:1 units:<id:5577006791947779410 details:<type_url:"type.googleapis.com/pb.v1.Bar" value:"\n\017the bar is open" > > units:<id:8674665223082153551 details:<type_url:"type.googleapis.com/pb.v1.Bar" value:"\n\021the bar is closed" > >
(*v1.Bar)(0xc0000694d0)(error:"the bar is open" )
(*v1.Bar)(0xc000069710)(error:"the bar is closed" )

--------------------------------------------------------------------------------
v1_bar_v2_foo.pb.bin
--------------------------------------------------------------------------------
(pb.Payload)value:4 units:<id:6334824724549167320 details:<type_url:"type.googleapis.com/pb.v1.Foo" value:"\n\016this is foo v1" > > units:<id:605394647632969758 details:<type_url:"type.googleapis.com/pb.v2.Foo" value:"\n\016this is foo v2\021\000\000\000\000\000\000\360?" > >
(*v1.Foo)(0xc000069920)(message:"this is foo v1" )
(*v2.Foo)(0xc000024ac0)(message:"this is foo v1" )
(*v1.Foo)(0xc000069d70)(message:"this is foo v2" 2:4607182418800017408 )
(*v2.Foo)(0xc000024b80)(message:"this is foo v2" ratio:1 )

--------------------------------------------------------------------------------
v1_foo.pb.bin
--------------------------------------------------------------------------------
(pb.Payload)value:2 units:<id:6129484611666145821 details:<type_url:"type.googleapis.com/pb.v1.Foo" value:"\n\016this is foo v1" > >
(*v1.Foo)(0xc000069fb0)(message:"this is foo v1" )
(*v2.Foo)(0xc000024d80)(message:"this is foo v1" )

--------------------------------------------------------------------------------
v1_foo_v1_bar.pb.bin
--------------------------------------------------------------------------------
(pb.Payload)value:3 units:<id:4037200794235010051 details:<type_url:"type.googleapis.com/pb.v1.Foo" value:"\n\016this is foo v1" > > units:<id:3916589616287113937 details:<type_url:"type.googleapis.com/pb.v1.Bar" value:"\n\021the bar is closed" > >
(*v1.Foo)(0xc0001ac240)(message:"this is foo v1" )
(*v2.Foo)(0xc000024f80)(message:"this is foo v1" )
(*v1.Bar)(0xc0001ac330)(error:"the bar is closed" )
```
