package main

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"net"
	"os"

	protos "github.com/bongerka/microservice/currency/protos/currency"
	"github.com/bongerka/microservice/currency/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	hlog := hclog.Default()

	gs := grpc.NewServer()
	c := server.NewCurrency(hlog)

	protos.RegisterCurrencyServer(gs, c)
	reflection.Register(gs)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 9092))
	if err != nil {
		hlog.Error("Unable to create listener", "error", err)
		os.Exit(1)
	}

	err = gs.Serve(l)
	hlog.Error("Unable to serve", "error", err)
}
