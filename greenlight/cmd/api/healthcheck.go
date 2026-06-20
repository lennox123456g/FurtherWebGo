package main

import (
	"fmt"
	"net/http"
)

//Declare a handler which writes plain info about status , operating e
//environment and version

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}
