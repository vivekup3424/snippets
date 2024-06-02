package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type requestBod struct {
	title   string
	content string
	expires int
}

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		//http.NotFound(w, r)
		w.Write([]byte("Error 404: Page not found"))
		app.InfoLogger.Println(w.Header())
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
		app.ErrorLogger.Println(err)
		http.Error(w, "internal server error:- failure to parse html", http.StatusInternalServerError)
		return
	}
	//render the parsed html pages on the endpoint
	if err := ts.Execute(w, nil); err != nil {
		app.ErrorLogger.Println(err)
		http.Error(w, "Internal server error when rendering html pages", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	//w.Write([]byte("Hello from Snippetbox\n"))
	app.InfoLogger.Println(w.Header())
}

// Add a snippetView handler function.
func (app *Application) snippetView(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Display a specific snippet...\n"))
	app.InfoLogger.Println(w.Header())
}
func (app *Application) snippetViewById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	//idNum := strconv.Atoi(id)
	str := fmt.Sprintf("Displaying the snippets of user with id = %s\n", id)
	idNum, _ := strconv.Atoi(id)
	snippets, err := app.snippets.Get(idNum)
	if err != nil {
		http.Error(w, "Bad shit happened", http.StatusBadGateway)
	}
	w.Write([]byte(str))
	fmt.Fprintf(w, "%+v", snippets)
	app.InfoLogger.Println(w.Header())
}

// Add a snippetCreate handler function.
func (app *Application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Create some variables holding dummy data. We'll remove these later on
	// during the build.
	var body []byte
	r.Body.Read(body)
	v := new(requestBod)
	json.Unmarshal(body, v)
	log.Println(body)
	w.Write([]byte(fmt.Sprintf("title = %s\n", v.content)))
	// Pass the data to the SnippetModel.Insert() method, receiving the
	// ID of the new record back.
	//id, err := app.snippets.Insert(title, content, expires)
	//if err != nil {
	//	app.ErrorLogger.Println(err)
	//	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	//	return
	//}
	//
	//// Redirect to the snippet view page.
	//http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
	//
	// Set response header and write response body.
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Create a new snippet...\n"))
	app.InfoLogger.Println(w.Header())
}
