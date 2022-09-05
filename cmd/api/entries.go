//Filename /cmd/api/entries.go

package main

import (
	"fmt"
	"net/http"
)

// Create Entry Handler - POST - /v1/entries
func (app *application) createEntryHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "create")

}

// get entries by Id - GET - /v1/entries/:id
func (app *application) showEntryHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	//display school id
	fmt.Fprintf(w, "Entries  %d\n", id)

}
