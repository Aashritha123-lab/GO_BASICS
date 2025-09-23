package main

import (
	"car/config"
	"car/handlers"
	"car/middleware"
	"fmt"
	"net/http"
)

func main() {

	config.ConnectDB()
	mux := http.NewServeMux()

	mux.HandleFunc("/cars", handlers.Carhandler)
	mux.HandleFunc("/cars/", handlers.Carhandler)
	wrappedmux := middleware.Logger(mux)
	wrappedmux = middleware.SecurityHeader(wrappedmux)

	fmt.Println("HTTP server listening....")
	http.ListenAndServe(":3051", wrappedmux)

}
