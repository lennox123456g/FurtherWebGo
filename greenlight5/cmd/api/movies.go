package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	//"strconv"
	//"github.com/julienschmidt/httprouter"
	"time"

	"greenlight.lennoxmugumira.net/internal/data" // New import
)

/*
// changing this

	func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "create a new Movie")
	}
*/
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	// Declare an anonymous struct to hold the information that we expect to be in the
	// HTTP request body (note that the field names and types in the struct are a subset
	// of the Movie struct that we created earlier). This struct will be our *target
	// decode destination*.
	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}

	// Initialize a new json.Decoder instance which reads from the request body, and
	// then use the Decode() method to decode the body contents into the input struct.
	// Importantly, notice that when we call Decode() we pass a *pointer* to the input
	// struct as the target decode destination. If there was an error during decoding,
	// we also use our generic errorResponse() helper to send the client a 400 Bad
	// Request response containing the error message.
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
	}

	// Dump the contents of the input struct in a HTTP response.
	fmt.Fprintf(w, "%+v\n", input)
}

// Add a showMovieHandler for the "GET /v1/movies/:id" endpoint. For now, we retrieve
// the interpolated "id" parameter from the current URL and include it in a placeholder
// response.

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		//http.NotFound(w, r)

		// Use the new notFoundResponse() helper.
		app.notFoundResponse(w, r)
		return
	}

	//encoding  a new Instance of the movie struct from internals/data
	//Create a new instance of the movie struct ,containing the ID we extracted from
	//the URL and some dummy Data. Note hat the year field we havent et it here
	movie := data.Movie{ //Data is the package mean , this means ,Movie from the data package
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romace", "war"},
		Version:   1,
	}

	//Create an envelope{"movie": movie } instance and pass it to writeJSON(), instead
	//of passing the plain movie struct
	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.logger.Print(err)
		//http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		// Use the new serverErrorResponse() helper.
		app.serverErrorResponse(w, r, err)
	}
}
