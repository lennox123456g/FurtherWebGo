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

	//USING MARSHALL
	//create a map thats ed the iformation that we want to send in the respojse

	/*
		data := map[string]string{
			"status ":     "available",
			"environment": app.config.env,
			"version":     version,
		}
	*/
	//Declare an envelope map containng the data for the response. Notice thata the way
	//we've constructed this means the environment and version data will now be nested
	//under a system_info key in the JSON response
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
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}

}
