package main

import (
	"net/http"

	"github.com/EmotionlessDev/todo-tasks/internal/data"
	"github.com/EmotionlessDev/todo-tasks/internal/validator"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello, world!"))
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// TODO: Implement the handlers for the task and list endpoints
func (app *application) taskHandlerGet(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	task, err := app.models.Task.Get(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"task": task}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// TODO: Implement the taskHandlerPatch handler
func (app *application) taskHandlerPost(w http.ResponseWriter, r *http.Request) {
	task := &data.Task{}
	err := app.readJSON(r, task)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	err = app.models.Task.Insert(task.Title, task.Description, task.ListID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusCreated, envelope{"task": task}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listHandlerGet(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	list, err := app.models.List.Get(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"list": list}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listHandlerPost(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json:"title"`
	}

	err := app.readJSON(r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	list := &data.List{
		Title: input.Title,
	}

	v := validator.New()
	if data.ValidateList(v, list); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.List.Insert(list)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// TODO: why this not working?
	// headers := make(http.Header)
	// headers.Set("Location", fmt.Sprintf("/v1/list/%d", list.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"list": list}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
