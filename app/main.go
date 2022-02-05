package main

import (
	"fmt"
	"log"
	"net/http"

	"./router"
	"github.com/ranajahanzaib/wdx"
)

func main() {
	wdx.Log("Hello World\n")
	r := router.Router()
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
