package main

import (
	"io"
	"net/http"
	"os"
	"os/signal"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	go start()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}

func start() {

	port := 8000
	router := mux.NewRouter()
	s1 := router.Host("www.domain.de").Path("/hallo").Subrouter()
	s1.Methods("GET").HandlerFunc(handler)
	// s1.Methods("POST").HandlerFunc(postHandler)
	http.ListenAndServe(":"+strconv.Itoa(port), router)
}

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!!\n")
}
