package main

import (
	"fmt"
	"net/http"

	//"strconv"
	//"github.com/julienschmidt/httprouter"
	"time"

	"greenlight.lennoxmugumira.net/internal/data" // New import
)

// Add a createMovieHandler for the "POST /v1/movies" endpoint. For now we simply
// return a plain-text placeholder response.
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new Movie")
}

// Add a showMovieHandler for the "GET /v1/movies/:id" endpoint. For now, we retrieve
// the interpolated "id" parameter from the current URL and include it in a placeholder
// response.

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
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
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
