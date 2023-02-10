package main

import (
	"log"
	"os"

	"github.com/alecthomas/repr"
	"github.com/min-organization/gosmi/parser"
)

func main() {
	module, err := parser.ParseFile(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	_ = module
	repr.Println(module)
}
