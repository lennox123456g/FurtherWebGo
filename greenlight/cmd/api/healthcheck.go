package main

import (
	//"fmt"
	"encoding/json" //New import
	"net/http"
)

//Declare a handler which writes plain info about status , operating e
//environment and version

// THIS HANDLER IMPLEMENTED AS A EMTHOD OF THE APPLICATION STRUCTE
// This is an effective and idiomatic way to make dependencies available to our handlers
// without resorting to global variables or closures — any dependency that the
// healthcheckHandler needs can simply be included as a field in the application struct when
// we initialize it in main()
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	/* BEFORE USING JSON VERSION
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
	*/

	// Create a fixed-format JSON response from a string. Notice how we're using a raw
	// string literal (enclosed with backticks) so that we can include double-quote
	// characters in the JSON without needing to escape them? We also use the %q verb to
	// wrap the interpolated values in double-quotes.

	/* before using  json.Marshal()

	js := `{"status": "available", "environment": %q, "version": %q}`
	js = fmt.Sprintf(js, app.config.env, version)
	*/

	//USING MARSHALL
	//create a map thats ed the iformation that we want to send in the respojse

	data := map[string]string{
		"status ":     "available",
		"environment": app.config.env,
		"version":     version,
	}

	//Phen pass the map to Marshall function which changes it to a []byte slice
	//containing encoded json , in case of error, we log it and send the clent a generic error message
	js, err := json.Marshal(data)
	if err != nil {
		app.logger.Print(err)
		http.Error(w, "Yhe server encountered a problem and couldnot process your request", http.StatusInternalServerError)
		return

	}

	// Append a newline to the JSON. This is just a small nicety to make it easier to
	// view in terminal applications.
	js = append(js, '\n')

	// Set the "Content-Type: application/json" header on the response. If you forget to
	// this, Go will default to sending a "Content-Type: text/plain; charset=utf-8"
	// header instead.
	w.Header().Set("Content-Type", "application/json")

	//Write the JSON as the HTTP response body
	//w.Write([]byte(js)) used besifre using Marchal fnction

	// Use w.Write() to send the []byte slice containing the JSON as the response body.
	w.Write(js)
}
