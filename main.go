package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/GodlikePenguin/GoServe/DirHandler"
	"github.com/GodlikePenguin/GoServe/FileHandler"
	"github.com/GodlikePenguin/GoServe/Utils"
)

func main() {
	http.HandleFunc("/", CustomFileServer)

	log.Printf("Serving on HTTP port: %s\n", "8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func CustomFileServer(w http.ResponseWriter, r *http.Request) {
	rootDir := os.Getenv("SERVE_LOCATION")
	if !strings.HasSuffix(rootDir, "/") {
		rootDir += "/"
	}
	currentItem := rootDir + strings.TrimPrefix(r.URL.Path, "/")
	log.Printf(currentItem)
	item, err := os.Stat(currentItem)
	if err != nil {
		log.Printf("Not a file or directory: %s", currentItem)
		Utils.ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	switch mode := item.Mode(); {
	case mode.IsDir():
		DirHandler.DirHandler(w, r, currentItem, r.URL.Path)
	case mode.IsRegular():
		FileHandler.FileHandler(w, r, currentItem)
	}
}
