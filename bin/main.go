package main

import (
	"flag"
	"log"
	"os"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/delaemon/go-thrift/common"
)

var (
	host   string
	port   string
	server bool
	client bool
)

func init() {
	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	f.StringVar(&host, "host", "0.0.0.0", "run client")
	f.StringVar(&port, "port", "8000", "run client")
	f.BoolVar(&client, "client", false, "run client")
	f.BoolVar(&server, "server", false, "run server")
	f.Parse(os.Args[1:])
	for 0 < f.NArg() {
		f.Parse(f.Args()[1:])
	}
	if !client && !server {
		f.PrintDefaults()
		log.Fatalln("too few arguments.")
	}
}

func main() {
	var transportFactory thrift.TTransportFactory = thrift.NewTTransportFactory()
	var protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
	var secure bool = false
	var addr = host + ":" + port

	if client {
		if err := common.RunClient(transportFactory, protocolFactory, addr, secure); err != nil {
			log.Fatalf("error running client: %s\n", err)
		}
	} else if server {
		if err := common.RunServer(transportFactory, protocolFactory, addr, secure); err != nil {
			log.Fatalf("error running server: %s\n", err)
		}
	} else {
		panic("failed launch.")
	}
}
