package main

import "net/http"

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello, world!"))
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
