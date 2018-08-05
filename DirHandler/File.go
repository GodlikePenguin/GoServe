package DirHandler

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type File struct {
	Name         string
	LastModified string
	Size         string
	Location     string
}

func CreateFile(base os.FileInfo, root string) *File {
	var size string
	if base.IsDir() {
		size = "-"
	} else {
		size = strconv.FormatInt(base.Size(), 10)
	}
	return &File{
		Name:         base.Name(),
		LastModified: base.ModTime().Format(time.UnixDate),
		Size:         size,
		Location:     fmt.Sprintf("%s/%s", strings.TrimSuffix(root, "/"), base.Name()),
	}
}
