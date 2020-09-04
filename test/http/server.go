package main

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

type server struct{}

func (s *server) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	body, _ := ioutil.ReadAll(req.Body)
	log.Printf("request body: %s", body)

	// remove User-Agent header
	// add custom User-Agent header to the request
	r := []byte(`{"3":["user-agent"],"1":{"user-agent":"Middleware!"}}`)

	rw.Write(r)
}

func main() {
	address := "127.0.0.1:3000"

	log.Printf("listen %s", address)

	ln, err := net.Listen("tcp", address)
	if err != nil {
		log.Printf("error listen, %v", err)
		os.Exit(1)
	}

	err = http.Serve(ln, &server{})
	if err != nil {
		log.Printf("error serve %s", err)
		os.Exit(1)
	}
}
