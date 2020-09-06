package main

import (
	traefik_plugin_bridge "github.com/negasus/traefik-plugin-bridge"
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

	resp := traefik_plugin_bridge.Response{
		DeleteRequestHeaders: []string{"user-agent"},
		AddRequestHeaders:    map[string]string{"user-agent": "Middleware!"},
	}

	respStr, _ := resp.ToJSON()

	rw.Write(respStr)
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
