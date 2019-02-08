package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func main() {
	StartBackendEphimeral()
}

var started = false

func StartBackendEphimeral() {
	if started {
		return
	}
	started = true
	http.HandleFunc("/", startPageHandler)
	listener, err := net.Listen("tcp", ":0")
	SocketAddress = fmt.Sprintf("%v", listener.Addr())

	_, _ = fmt.Fprintf(os.Stderr, "Running on port "+SocketAddress)

	if err = http.Serve(listener, nil); err != nil {
		panic(err)
	}
}

var SocketAddress string
