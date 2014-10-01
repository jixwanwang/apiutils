package apiutils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func RequireParams(form url.Values, params []string) error {
	for _, param := range params {
		if len(form[param]) == 0 {
			return fmt.Errorf("Missing param: %s", param)
		}
	}
	return nil
}

// ReadParams reads in parameters from the request, using the content type.
func ReadParams(r *http.Request) (map[string]interface{}, error) {
	params := make(map[string]interface{})

	if r.Header.Get("Content-Type") == "application/json" {
		decoder := json.NewDecoder(r.Body)
		return params, decoder.Decode(&params)
	} else {
		r.ParseForm()

		// Take first argument, equivalent to Get()
		for k, v := range r.Form {
			params[k] = v[0]
		}

		return params, nil
	}
}

func RequireFormParams(r *http.Request, params []string) error {
	for _, param := range params {
		if len(r.FormValue(param)) == 0 {
			return fmt.Errorf("Missing param: %s", param)
		}
	}
	return nil
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func NewErrorResponse(status int, message string) ErrorResponse {
	return ErrorResponse{
		Status:  status,
		Message: message,
		Error:   http.StatusText(status),
	}
}

func ServeJSON(w http.ResponseWriter, v interface{}) {
	content, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Length", strconv.Itoa(len(content)))
	w.Header().Set("Content-Type", "application/json")
	w.Write(content)
}

func ServeError(w http.ResponseWriter, errRes ErrorResponse) {
	w.WriteHeader(errRes.Status)
	ServeJSON(w, errRes)
}
