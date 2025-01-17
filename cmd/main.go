package main

import (
	"github.com/linuxerwang/plz-exercise/proto/data"
	"google.golang.org/protobuf/encoding/prototext"
)

func main() {
	p := data.Product{
		Name: "jacket",
	}
	prototext.Marshal(&p)
}
