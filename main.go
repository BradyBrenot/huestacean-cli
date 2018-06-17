package main

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

	pb "github.com/BradyBrenot/huestacean-cli/proto-gen"
	"github.com/jessevdk/go-flags"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

//Options contains user-provided command-line settings
var Options struct {
	HostAddress string `short:"h" long:"host" description:"Host address to use to connect to huestaceand" default:"localhost:55610"`
}

func connect() (*grpc.ClientConn, pb.HuestaceanServerClient, context.Context, context.CancelFunc) {
	conn, err := grpc.Dial(Options.HostAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial failed %v", err)
	}
	c := pb.NewHuestaceanServerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	return conn, c, ctx, cancel
}

//////////////////////////////////////////////////////////////////////////

type getProvidersCommand struct {
}

//Execute for getProvidersCommand
func (cmd *getProvidersCommand) Execute(args []string) error {
	conn, client, ctx, cancel := connect()
	defer conn.Close()
	defer cancel()

	r, err := client.GetDeviceProviders(ctx, &pb.GetDeviceProvidersRequest{})
	if err != nil {
		return fmt.Errorf("GetDeviceProviders failed: %v", err)
	}

	log.Println("Providers:")

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 5, 0, 2, ' ', 0)
	fmt.Fprintln(w, "id\tname\tarchetype\tstate")
	for key, provider := range r.GetProviders() {
		fmt.Fprintf(w, "%v\t%v\t%v\t%v", key, provider.Name, provider.ArchetypeId, provider.State)
	}
	w.Flush()

	return nil
}

//////////////////////////////////////////////////////////////////////////

type linkCommand struct {
	providerID uint32 `short:"p" long:"provider" description:"ID of provider to link. Use getDeviceProviders to list." required:"true"`
}

//Execute for getProvidersCommand
func (cmd *linkCommand) Execute(args []string) error {
	conn, client, ctx, cancel := connect()
	defer conn.Close()
	defer cancel()

	_, err := client.Link(ctx, &pb.LinkRequest{
		DeviceId: cmd.providerID,
	})
	if err != nil {
		return fmt.Errorf("Link failed: %v", err)
	}

	fmt.Println("Link request sent")

	return nil
}

//////////////////////////////////////////////////////////////////////////

var parser = flags.NewParser(&Options, flags.Default)

func main() {
	parser.AddCommand("getProviders", "get the list of currently-known device providers (e.g. Hue bridges)", "", &getProvidersCommand{})
	parser.AddCommand("link", "pair the given device provider. You should press the link button on the hardware before sending this command", "", &linkCommand{})
	parser.Parse()
}
