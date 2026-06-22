package main

import (
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
