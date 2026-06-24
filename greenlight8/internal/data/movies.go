package data

import (
	"time"

	"greenlight.lennoxmugumira.net/internal/validator" // New import
)

// Annotate the Movie struct with struct tags to control how the keys appear in the
// JSON-encoded output.
type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Year      int32     `json:"year"`
	//Runtime   int32     `json:"runtime"`

	//ADVANCED CUSTOMIZATION
	// Use the Runtime type instead of int32. Note that the omitempty directive will
	// still work on this: if the Runtime field has the underlying value 0, then it will
	// be considered empty and omitted -- and the MarshalJSON() method we just made
	// won't be called at all.
	Runtime Runtime  `json:"runtime,omitempty"`
	Genres  []string `json:"genres"`
	Version int32    `json:"version"`
}

/*
OMIT EMPTY AND - STRUCT DIRECTIVES


type Movie struct {
    ID        int64     `json:"id"`
    CreatedAt time.Time `json:"-"` // Use the - directive
    Title     string    `json:"title"`
    Year      int32     `json:"year,omitempty"`    // Add the omitempty directive
    Runtime   int32     `json:"runtime,omitempty"` // Add the omitempty directive
    Genres    []string  `json:"genres,omitempty"`  // Add the omitempty directive
    Version   int32     `json:"version"`
}
*/

func ValidateMovie(v *validator.Validator, movie *Movie) {
	v.Check(movie.Title != "", "title", "must be provided")
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year >= 1888, "year", "must be greater than 1888")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")
	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")
	v.Check(movie.Genres != nil, "genres", "must be provided")
	v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate values")
} //we made a custom function for the vaildaors to the actual movies struct that lives

//In large projects it’s likely that you’ll want to reuse some of the same validation checks in
//multiple places. In our case — for example — we’ll want to use many of these same checks
//later when a client edits the movie data.
//To prevent duplication, we can collect the validation checks for a movie into a standalone
//ValidateMovie() function. In theory this function could live almost anywhere in our
//codebase — next to the handlers in the cmd/api/movies.go file, or possibly in the
//internal/validators package. But personally, I like to keep the validation checks close to
//the relevant domain type in the internal/data package.

//important: Notice that the validation checks are now being performed on a Movie
//struct — not on the input struct in our handlers.
