package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/golang/protobuf/proto"

	"github.com/j-vizcaino/pb-test/pb"
	v1 "github.com/j-vizcaino/pb-test/pb/v1"
	v2 "github.com/j-vizcaino/pb-test/pb/v2"
)

func printPayload(filename string) {
	m := pb.Payload{}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("%s: %e\n", filename, err)
		return
	}

	err = proto.Unmarshal(data, &m)
	if err != nil {
		fmt.Printf("%s: Protobuf unmarshal failed, %e", filename, err)
		return
	}
	ruler := strings.Repeat("-", 80)
	spew.Printf("%s\n%s\n%s\n%#+v\n", ruler, filename, ruler, m)
	defer fmt.Println()

	if len(m.Units) == 0 {
		return
	}

	for _, unit := range m.Units {
		switch unit.Details.TypeUrl {
		case "type.googleapis.com/pb.v1.Bar":
			spew.Dump(unmarshalDetails(unit.Details.Value, &v1.Bar{}))
		case "type.googleapis.com/pb.v1.Foo":
			fallthrough
		case "type.googleapis.com/pb.v2.Foo":
			spew.Dump(unmarshalDetails(unit.Details.Value, &v1.Foo{}))
			spew.Dump(unmarshalDetails(unit.Details.Value, &v2.Foo{}))
		}
	}
}

func unmarshalDetails(d []byte, out proto.Message) interface{} {
	err := proto.Unmarshal(d, out)
	if err != nil {
		return fmt.Sprintf("Failed to unmarshal details to %T, %e", out, err)
	}
	return out
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: invalid number of arguments")
		fmt.Println("Usage: reader file.pb.bin [...]")
		os.Exit(1)
	}

	for i := 1; i < len(os.Args); i++ {
		printPayload(os.Args[i])
	}
}
