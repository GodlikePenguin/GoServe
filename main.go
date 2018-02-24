package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"fmt"

	"github.com/GodlikePenguin/GoServe/DirHandler"
	"github.com/GodlikePenguin/GoServe/Utils"
	"github.com/GodlikePenguin/GoServe/FileHandler"
)

func main() {
	someString := "/Program Files"
	fmt.Println(strings.LastIndex(someString, "/"))
	fmt.Println(someString[:strings.LastIndex(someString, "/")])
	//os.Exit(0)

	http.HandleFunc("/", CustomFileServer)

	log.Printf("Serving on HTTP port: %s\n", "8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func CustomFileServer(w http.ResponseWriter, r *http.Request) {
	rootDir := "D:/"
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
		DirHandler.DirHandler(w, r, currentItem)
	case mode.IsRegular():
		FileHandler.FileHandler(w, r, currentItem)
	}
}
