package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//Define the route handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world!") //Send "Hello world!"as the response
	}
	//Create a new server mux
	mux := http.NewServeMux()
	//Register the route handler function to handle all requests to "/"
	mux.HandleFunc("/", handler)
	//Create a new http server
	server := &http.Server{
		Addr:    ":8080", //listen on port 8080
		Handler: mux,
	}
	//start the server
	log.Println("Server listening on port 8080...")
	log.Fatal(server.ListenAndServe())
}
