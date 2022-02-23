package main

import (
	"log"
	"net/http"
	"os"

	"github.com/adityagoel/product-api/handlers"
)

func main() {
	thisLogger := log.New(os.Stdout, "product-api", log.LstdFlags)

	thisHelloHandler := handlers.NewHello(thisLogger)
	thisGoodByeHandler := handlers.NewGoodBye(thisLogger)

	thisServeMux := http.NewServeMux()

	thisServeMux.Handle("/", thisHelloHandler)
	thisServeMux.Handle("/goodBye", thisGoodByeHandler)

	http.ListenAndServe(":8081", thisServeMux)
}
