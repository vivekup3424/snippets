package main

import (
	"log"
	"net/http"
)

func main() {
	// Register the two new handler functions and corresponding URL patterns with
	// the servemux, in exactly the same way that we did before.
	router := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	router.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	//Registering the routes
	router.HandleFunc("GET /", home)
	router.HandleFunc("GET /snippet/view", snippetView)
	router.HandleFunc("GET /snippet/view/{id}", snippetViewById)
	router.HandleFunc("POST /snippet/create", snippetCreate)
	log.Println("Starting server on 127.0.0.1:4001")
	//making the server listen on appropriate port
	if err := http.ListenAndServe("localhost:4001", router); err != nil {
		log.Fatal(err)
	}
}
