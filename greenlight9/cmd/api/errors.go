package main

import (
	"fmt"
	"net/http"
)

// The logError() method is a generic helper for logging an error message.
// we shall use a structured logger later
func (app *application) logError(r *http.Request, err error) {
	app.logger.Print(err)
}

// errorResponse() method id  ageneric helpwer for  sending JSON formtted error
// messags to client with a given status code Note that we're using an any
// type for the message parameter, rather than just a string type, as this gives us
// more flexibility over the values that we can include in the response.
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"err": message}

	//write response  usin WriteJSON() helper. If this happens to return an
	//error then log it, and all back to the sending c lientand empt response with a
	//500 internal server Error status code
	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

// The serverErrorResponse() method will be used when our application encounters an
// unexpected problem at runtime. It logs the detailed error message, then uses the
// errorResponse() helper to send a 500 Internal Server Error status code and JSON
// response (containing a generic error message) to the client.
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// The notFpundResponse() method wll be used to send  a 404 No found status code and
// JSON Response to the cliet
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

// The methodNotAllowedResponse() method will be used to send a 405 Method Not Allowed
// status code and JSON response to the client.
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

//Now those are in place, let’s update our API handlers to use these new helpers instead of
//the http.Error() and http.NotFound() functions

// VALIDITY
// Note that the errors parameter here has the type map[string]string, which is exactly
// the same as the errors map contained in our Validator type.
func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	app.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
} //Then once that’s done, head back to your createMovieHandler and update it to perform the
//necessary validation checks on the input struct. Like so

// The badRequestResponse() method will be used to send a 400 Bad Request status code
// and JSON response to the client, including the error message from the err parameter.
func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}
