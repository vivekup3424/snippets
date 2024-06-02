package main

import (
	"log"
	"net/http"
	"os"
)

// Use log.New() to create a logger for writing information messages. This takes
// three parameters: the destination to write the logs to (os.Stdout), a string
// prefix for message (INFO followed by a tab), and flags to indicate what
// additional information to include (local date and time). Note that the flags
// are joined using the bitwise OR operator |.
var InfoLogger = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

// Create a logger for writing error messages in the same way, but use stderr as
// the destination and use the log.Lshortfile flag to include the relevant
// file name and line number.
var ErrorLogger = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

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
	InfoLogger.Println("Starting server on 127.0.0.1:4001")
	//making the server listen on appropriate port
	if err := http.ListenAndServe("localhost:4001", router); err != nil {
		ErrorLogger.Fatal(err)
	}
}
