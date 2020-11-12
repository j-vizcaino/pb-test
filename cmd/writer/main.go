package main

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"

	"github.com/j-vizcaino/pb-test/pb"
	v1 "github.com/j-vizcaino/pb-test/pb/v1"
	v2 "github.com/j-vizcaino/pb-test/pb/v2"
)

func writePayload(p *pb.Payload, out string) error {
	data, err := proto.Marshal(p)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(out + ".pb.bin", data, 0644)
}


func newPayloadWithDetails(val int, details proto.Message) *pb.Payload {
	p := &pb.Payload{
		Value: int32(val),
	}
	if details == nil {
		return p
	}
	any, err := ptypes.MarshalAny(details)
	if err != nil {
		panic(err)
	}
	p.Details = any
	return p
}

func main() {
	outFiles := map[string]*pb.Payload{
		"no_details": newPayloadWithDetails(42, nil),
		"v1_bar": newPayloadWithDetails(1, &v1.Bar{Error: "the bar is closed"}),
		"v1_foo": newPayloadWithDetails(1, &v1.Foo{Message: "this is foo v1"}),
		"v2_foo": newPayloadWithDetails(2, &v2.Foo{Message: "this is foo v2", Ratio: 1.0}),
	}

	for name, message := range outFiles {
		err := writePayload(message, name)
		if err != nil {
			fmt.Printf("Failed to create %s: %e\n", name, err)
		} else {
			fmt.Println("Wrote", name)
		}
	}
}

