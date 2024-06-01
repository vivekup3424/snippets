package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		//http.NotFound(w, r)
		w.Write([]byte("Error 404: Page not found"))
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from Snippetbox\n"))
}

// Add a snippetView handler function.
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Display a specific snippet...\n"))
}
func snippetViewById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	//idNum := strconv.Atoi(id)
	str := fmt.Sprintf("Displaying the snippets of user with id = %s\n", id)
	w.Write([]byte(str))
}

// Add a snippetCreate handler function.
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	//I dont need this since Go 1.22 handles this itself
	//if r.Method != "POST" {
	// 	w.Header().Set("Allow","POST")
	//	w.WriteHeader(405)
	//	w.Write([]byte("Method Not Allowed"))
	//
	//}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Create a new snippet...\n"))
}
