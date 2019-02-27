package protorpc

import (
	"context"
	"time"

	"github.com/mmmckay/grpc-v-http2/protorpc/routeguide"
	"google.golang.org/grpc"
)

// Send sends with a stream
func Send(b []byte, count int) (time.Duration, error) {
	var err error
	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	client := routeguide.NewRouteGuideClient(conn)

	stream, err := client.GetHTML(context.Background())
	if err != nil {
		return 0, err
	}
	start := time.Now()
	for i := 0; i < count; i++ {
		if err = stream.Send(&routeguide.Doc{HTML: b, Collection: "stuff"}); err != nil {
			break
		}
	}
	if err != nil {
		return 0, err
	}
	_, err = stream.CloseAndRecv()
	if err != nil {
		return 0, err
	}
	return time.Now().Sub(start), nil

}
