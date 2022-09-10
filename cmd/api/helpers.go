//Filename:cmd/api/helpers.go
package main

import	(
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readIDParam (r *http.Request) (int64, error){
	//use the '*ParamsfromContext())' function to get the request context as a slice
	params := httprouter.ParamsFromContext(r.Context())
	//get the value of the 'id' parameter
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0 , errors.New("Invalid ID Parameter")
	}
	return id, nil
}