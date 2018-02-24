package Utils

import (
	"net/http"
	"fmt"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	switch status {
	case http.StatusNotFound:
		fmt.Fprint(w, "<h1>404 File Not Found</h1>")
	case http.StatusNotImplemented:
		fmt.Fprint(w, "<h1>501 Not Implemented</h1>")
	case http.StatusInternalServerError:
		fmt.Fprint(w, "<h1>500 Internal Server Error</h1>")
	}
}
