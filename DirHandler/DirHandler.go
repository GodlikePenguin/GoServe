package DirHandler

import (
	"net/http"
	"io/ioutil"
	"github.com/GodlikePenguin/GoServe/Utils"
	"log"
)

func DirHandler(w http.ResponseWriter, r *http.Request, directory string) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Printf(err.Error())
		Utils.ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	structure := Directory{name:directory, fileList: files}
	w.Write(structure.createHttpStructure())
}
