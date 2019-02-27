package protorpc

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	rg "github.com/mmmckay/grpc-v-http2/protorpc/routeguide"
	"google.golang.org/grpc"
)

var counter int

type server struct{}

func (s *server) GetHTML(stream rg.RouteGuide_GetHTMLServer) error {
	for {
		doc, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&rg.Result{Answer: "done"})
		}
		if err != nil {
			return err
		}
		_ = doc
		counter++
	}
}

func newServer() *server {
	return &server{}
}

// Serve serves
func Serve() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8081))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	rg.RegisterRouteGuideServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
