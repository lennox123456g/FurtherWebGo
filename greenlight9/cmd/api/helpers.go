package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"strings" //New Import

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

// NEW HELPER FUNCTIO TO HANDLE THE READING OF JSO AND ALSO CATER FOR THE ERRORS THAT RESULT FORM DECODE()
func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	//Use hhtp.MaxBytesReader() to limit the size of the request Body to 1MB
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	//Initialise the json.Decoder, and call the DisallowunKnownFileds method on it
	//befire decoding. This means that if the JSON from from the client now includes any
	//field which cannot be mapped to target destination, the decorder will return
	//An error instead of just ignoring the field
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	// Decode the request body into the target destination.
	err := dec.Decode(dst)
	if err != nil {
		// If there is an error during decoding, start the triage...
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError
		var maxBytesError *http.MaxBytesError

		switch {
		// Use the errors.As() function to check whether the error has the type
		// *json.SyntaxError. If it does, then return a plain-english error message
		// which includes the location of the problem.
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)
		// In some circumstances Decode() may also return an io.ErrUnexpectedEOF error
		// for syntax errors in the JSON. So we check for this using errors.Is() and
		// return a generic error message. There is an open issue regarding this at
		// https://github.com/golang/go/issues/25956.
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")

		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		//If the jSOn contains a field which cannot bre mapped to the traget destination
		//then Decode() will  will now return an error message in the format "json:unknown"
		//field "<name>" . We check for this, extract the field name from the error,
		// and interpolate it into our custom error message. Note that there's an open
		// issue at https://github.com/golang/go/issues/29035 regarding turning this
		// into a distinct error type in the future.
		case strings.HasPrefix(err.Error(), "json: unknown field"):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field")
			return fmt.Errorf("body contains unknown key %s", fieldName)

			// Use the errors.As() function to check whether the error has the type
		// *http.MaxBytesError. If it does, then it means the request body exceeded our
		// size limit of 1MB and we return a clear error message.
		case errors.As(err, &maxBytesError):
			return fmt.Errorf("body must not be larger than %d bytes", maxBytesError.Limit)

		case errors.As(err, &invalidUnmarshalError):
			panic(err)
		// For anything else, return the error message as-is.
		default:
			return err
		}
	}

	// Call Decode() again, using a pointer to an empty anonymous struct as the
	// destination. If the request body only contained a single JSON value this will
	// return an io.EOF error. So if we get anything else, we know that there is
	// additional data in the request body and we return our own custom error message.
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON value ")
	}
	return nil
} //Now go and updae createMovieHandler  in Movies.go
