package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		//http.NotFound(w, r)
		w.Write([]byte("Error 404: Page not found"))
		InfoLogger.Println(w.Header())
		return
	}
	// Use the template.ParseFiles() function to read the template file into a
	// template set. If there's an error, we log the detailed error message and use
	// the http.Error() function to send a generic 500 Internal Server Error
	// response to the user.
	path := "./ui/html/pages/home.html"
	ts, err := template.ParseFiles(path)
	//log the error as 500 on the system
	// We then use the Execute() method on the template set to write the
	// template content as the response body. The last parameter to Execute()
	// represents any dynamic data that we want to pass in, which for now we'll
	// leave as nil.
	if err != nil {
		ErrorLogger.Println(err)
		http.Error(w, "internal server error:- failure to parse html", http.StatusInternalServerError)
		return
	}
	//render the parsed html pages on the endpoint
	if err := ts.Execute(w, nil); err != nil {
		ErrorLogger.Println(err)
		http.Error(w, "Internal server error when rendering html pages", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	//w.Write([]byte("Hello from Snippetbox\n"))
	InfoLogger.Println(w.Header())
}

// Add a snippetView handler function.
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Display a specific snippet...\n"))
	InfoLogger.Println(w.Header())
}
func snippetViewById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	//idNum := strconv.Atoi(id)
	str := fmt.Sprintf("Displaying the snippets of user with id = %s\n", id)
	w.Write([]byte(str))
	InfoLogger.Println(w.Header())
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
	InfoLogger.Println(w.Header())
}
