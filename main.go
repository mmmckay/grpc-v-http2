package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mmmckay/grpc-v-http2/http2"
	"github.com/mmmckay/grpc-v-http2/protorpc"
)

func main() {
	b := RandStringBytesMaskImprSrc(50000)
	go http2.Serve()
	go protorpc.Serve()

	f, err := os.OpenFile("data.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	defer f.Close()
	if err != nil {
		panic(err.Error())
	}

	for i := 1; i < 100000; i += i {
		http2Duration, err := http2.Send(b, i)
		if err != nil {
			panic(err.Error())
		}

		protoDuration, err := protorpc.Send(b, i)
		if err != nil {
			panic(err.Error())
		}

		f.WriteString(fmt.Sprintf("%.4f,%.4f,%d,\n", float64(http2Duration)/float64(time.Second), float64(protoDuration)/float64(time.Second), i))
	}
}
