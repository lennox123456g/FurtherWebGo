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

// HELPER FUNCTION TO HELP IN JSON ENCODING  FOr  REUSE
// Define a writeJSON() helper for sending responses. This takes the destination
// http.ResponseWriter, the HTTP status code to send, the data to encode to JSON, and a
// header map containing any additional HTTP headers we want to include in the response.
func (app *application) writeJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	//Encode teh data to JSON , rteurning the eror i there isa any
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	//Append a newline to make it easier to view in terminakl applications
	js = append(js, '\n')

	// At this point, we know that we won't encounter any more errors before writing the
	// response, so it's safe to add any headers that we want to include. We loop
	// through the header map and add each header to the http.ResponseWriter header map.
	// Note that it's OK if the provided header map is nil. Go doesn't throw an error
	// if you try to range over (or generally, read from) a nil map.
	for key, value := range headers { //headers is a  map
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
