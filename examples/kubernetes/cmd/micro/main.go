package main

import (
	"os"

	"github.com/asim/go-micro/v3/broker"
	"github.com/asim/go-micro/v3/client"
	cli "github.com/asim/go-micro/v3/client/grpc"
	"github.com/asim/go-micro/v3/config/cmd"
	"github.com/asim/go-micro/v3/server"
	srv "github.com/asim/go-micro/v3/server/grpc"
	bkr "github.com/micro/go-plugins/broker/grpc/v2"
	_ "github.com/micro/go-plugins/registry/kubernetes/v2"

	// static selector offloads load balancing to k8s services
	// enable with MICRO_SELECTOR=static or --selector=static
	// requires user to create k8s services
	_ "github.com/micro/go-plugins/client/selector/static/v2"

	// disable namespace by default
	_ "github.com/asim/go-micro/v3/api"
)

func main() {

	// set values for registry/selector
	os.Setenv("MICRO_REGISTRY", "kubernetes")
	os.Setenv("MICRO_SELECTOR", "static")

	// setup broker/client/server
	broker.DefaultBroker = bkr.NewBroker()
	client.DefaultClient = cli.NewClient()
	server.DefaultServer = srv.NewServer()

	// init command
	cmd.Init()
}
