package response

import (
	"encoding/json"
	"net/http"
)

func respond(w http.ResponseWriter, code int, src interface{}) {
	var body []byte
	var err error

	switch s := src.(type) {
	case []byte:
		if !json.Valid(s) {
			BuildError(w, http.StatusInternalServerError, err)
			return
		}
		body = s
	case string:
		body = []byte(s)
	default:
		if body, err = json.Marshal(src); err != nil {
			BuildError(w, http.StatusInternalServerError, err)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	w.Write(body)
}

// BuildJSON is wrapped Respond when success response
func BuildJSON(w http.ResponseWriter, code int, src interface{}) {
	respond(w, code, src)
}

// BuildError is wrapped Respond when error response
func BuildError(w http.ResponseWriter, code int, err error) {
	respond(w, code, err)
}
