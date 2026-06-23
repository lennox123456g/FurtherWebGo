package main

import (
	//"fmt"
	//"encoding/json" //New import
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

	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, env, nil) //changed from data to env
	if err != nil {
		app.logger.Print(err)
		//http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError) //changing to use our errors.go file
		app.serverErrorResponse(w, r, err)
	}

}
