//Filename /cmd/api/healthcheck.go

package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	//create a map to hold our healthcheck data

	fmt.Fprintln(w, "Healthcheck")

}
