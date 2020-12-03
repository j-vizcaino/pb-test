package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"

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
	return ioutil.WriteFile(out+".pb.bin", data, 0644)
}

func newPayloadWithDetails(val int, messages ...proto.Message) *pb.Payload {
	p := &pb.Payload{
		Value: int32(val),
	}
	if len(messages) == 0 {
		return p
	}

	for _, message := range messages {
		any, err := ptypes.MarshalAny(message)
		if err != nil {
			panic(err)
		}
		p.Units = append(p.Units, &pb.Unit{
			Id:      rand.Int63(),
			Details: any,
		})
	}
	return p
}

func main() {
	outFiles := map[string]*pb.Payload{
		"no_details":    newPayloadWithDetails(42),
		"v1_bar":        newPayloadWithDetails(1, &v1.Bar{Error: "the bar is open"}, &v1.Bar{Error: "the bar is closed"}),
		"v1_foo":        newPayloadWithDetails(2, &v1.Foo{Message: "this is foo v1"}),
		"v1_foo_v1_bar": newPayloadWithDetails(3, &v1.Foo{Message: "this is foo v1"}, &v1.Bar{Error: "the bar is closed"}),
		"v1_bar_v2_foo": newPayloadWithDetails(4, &v1.Foo{Message: "this is foo v1"}, &v2.Foo{Message: "this is foo v2", Ratio: 1.0}),
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
