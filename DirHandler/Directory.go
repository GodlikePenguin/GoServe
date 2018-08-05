package DirHandler

import (
	"os"
)

type Directory struct {
	Name     string
	BackPath string
	Files    []*File
}

func CreateDirectory(name string, backPath string, files []os.FileInfo, root string) *Directory {
	dir := &Directory{
		Name:     name,
		BackPath: backPath,
	}
	var fileList []*File
	for _, file := range files {
		fileList = append(fileList, CreateFile(file, root))
	}
	dir.Files = fileList
	return dir
}
