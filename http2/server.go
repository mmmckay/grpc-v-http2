package http2

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var counter int

// Serve serves
func Serve() {
	r := mux.NewRouter()
	r.HandleFunc("/", Receive).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Doc for html
type Doc struct {
	HTML       []byte
	Collection string
}

// Receive handler
func Receive(w http.ResponseWriter, r *http.Request) {
	doc := &Doc{}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "bad data", http.StatusBadRequest)
	}
	err = json.Unmarshal(data, doc)
	_, _ = ioutil.ReadAll(r.Body)
	r.Body.Close()

	counter++
}
