package main

import (
	"encoding/json" //New import
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// HELPER FUNCTION FOR GETTING THE ID

// Retrieve the "id" URL from the current request context
// then convert it to an integer  and return it.If the operation isnt successul, return O and an errror
func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("Invalid Id Parameter")
	}

	return id, nil
}

// Deining the Envelope type
type envelope map[string]any

// Change the data parameter to have the type envelope instead of any.
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error { //changed the any to envelope
	// Use the json.MarshalIndent() function so that whitespace is added to the encoded
	// JSON. Here we use no line prefix ("") and tab indents ("\t") for each element.
	js, err := json.MarshalIndent(data, "", "\t") //change to use Marchal Indent function to make for spacing fpor json out put in terminak when we use curll
	if err != nil {
		return err
	}
	//Append a newline to make it easier to view in terminakl applications
	js = append(js, '\n')

	for key, value := range headers { //headers is a  map
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
