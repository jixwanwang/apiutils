package apiutils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

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

func RequireParams(form url.Values, params []string) error {
	for _, param := range params {
		if len(form[param]) == 0 {
			return fmt.Errorf("Missing param: %s", param)
		}
	}
	return nil
}
