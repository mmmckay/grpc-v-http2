package http2

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// Send sends http2 requests
func Send(b []byte, count int) (time.Duration, error) {
	d := &Doc{HTML: b, Collection: "stuff"}
	jsonBytes, err := json.Marshal(d)
	if err != nil {
		return 0, err
	}
	start := time.Now()
	for i := 0; i < count; i++ {
		resp, err := http.Post("http://localhost:8080/", "application/json", bytes.NewReader(jsonBytes))
		if err != nil {
			return 0, err
		}
		_, _ = ioutil.ReadAll(resp.Body)
		resp.Body.Close()
	}

	return time.Now().Sub(start), nil
}
