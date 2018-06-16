package main

import (
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/BradyBrenot/huestacean-cli/proto-gen"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	HostAddress string `short:"h" long:"host" description:"Host address to use to connect to huestaceand" default:"localhost:55610"`
}

var parser = flags.NewParser(&opts, flags.Default)

func main() {
	_, err := parser.Parse()
	if err != nil {
		return;
	}

	conn, err := grpc.Dial(opts.HostAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial failed %v", err)
	}
	defer conn.Close()
	c := pb.NewHuestaceanServerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetDeviceProviders(ctx, &pb.GetDeviceProvidersRequest{})
	if err != nil {
		log.Fatalf("GetDeviceProviders failed: %v", err)
	}
	log.Printf("Response: %v", r)
	log.Printf("Response map: %v", r.Providers)
}