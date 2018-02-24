package FileHandler

import (
	"net/http"
)

func FileHandler(w http.ResponseWriter, r *http.Request, path string) {
	http.ServeFile(w, r, path)
	return
}
