apiutils
========

Helpful API handling utilities.


## Getting Started

  go get github.com/jixwanwang/apiutils

## Example Usage:

```
package main

import (
        "log"
        "net/http"

        "github.com/jixwanwang/apiutils"
)

func main() {
        http.HandleFunc("/dowork", workwork)

        if err := http.ListenAndServe(":8080", http.DefaultServeMux); err != nil {
                log.Fatalf("Could not start server: %s", err.Error())
        }
}

func workwork(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		err := apiutils.RequireParams(query, []string{"foo", "bar"})
		if err != nil {
			apiutils.ServeError(w, apiutils.NewErrorResponse(http.StatusBadRequest, err.Error()))
			return
		}
        apiutils.ServeJSON(w, []string{"Hello!", "World"})
}
```
