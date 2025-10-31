package main

import (
	"fmt"
	"net/http"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
	"github.com/mdger/forever-wapp/controller"
)

func main() {
	// Start the server!
	router := controller.Init()
	addr := "127.0.0.1:8888"
	fmt.Println("starting server on", addr)
	_ = http.ListenAndServe(addr, router)
}
