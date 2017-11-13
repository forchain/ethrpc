package main

import (
	"github.com/forchain/ethrpc/parser"
	"flag"
	"os"
)

func main() {

	rpc := new(string)
	out := new(string)

	flag.StringVar(rpc, "rpc", "127.0.0.1:8545", "RPC server address")
	flag.StringVar(out, "out", "/tmp/rdf", "Output directory")
	flag.Parse()

	if _, err := os.Stat(*out); os.IsNotExist(err) {
		os.Mkdir(*out, 0666)
	}

	parser.Parse(*rpc, *out)
}
