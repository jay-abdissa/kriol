//Filename:cmd/api/healthcheck.go
package main

import	(
	"net/http"
)
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request)	{
	//create a map to hod our healthcheck data
	data := map[string]string{
		"status": "available",
		"environment": app.config.env,
		"version": version,
	}
	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil{
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}
	//specify that we will serve our responses using json
	//w.Header().Set("Content-Type", "application/json")
	//write the json as a HTTP response body
	//w.Write([]byte(js))

}
