package main

import (
	"encoding/json"
	"net/http"
)

type envelope map[string]interface{}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(js)
	if err != nil {
		return err
	}

	return nil
}

// func (app *application) readJSON(r *http.Request, dst interface{}) error {
// 	dec := json.NewDecoder(r.Body)
// 	dec.DisallowUnknownFields()
//
// 	err := dec.Decode(dst)
//
// 	if err != nil {
// 		var syntaxError *json.SyntaxError
// 		var unmarshalTypeError *json.UnmarshalTypeError
// 		var invalidUnmarshalError *json.InvalidUnmarshalError
// 		switch {
// 		case errors.As(err, &syntaxError):
// 			return fmt.Errorf("request body contains badly-formed JSON (at character %d)", syntaxError.Offset)
// 		case errors.Is(err, io.ErrUnexpectedEOF):
// 			return errors.New("request body contains badly-formed JSON")
// 		case errors.As(err, &unmarshalTypeError):
// 			return fmt.Errorf("request body contains an invalid value for the %q field (at character %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
// 		case errors.Is(err, io.EOF):
// 			return errors.New("request body must not be empty")
// 		case errors.As(err, &invalidUnmarshalError):
// 			panic(err)
// 		default:
// 			return err
// 		}
// 	}
//
// 	err = dec.Decode(&struct{}{})
// 	if err != io.EOF {
// 		return errors.New("request body must only contain a single JSON value")
// 	}
//
// 	return nil
// }
