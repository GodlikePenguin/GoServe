package DirHandler

import (
	"os"
	"strconv"
	"time"
	"strings"
	"github.com/GodlikePenguin/GoServe/Utils"
)

type Directory struct {
	name string
	fileList []os.FileInfo
}

func (d Directory) createHttpStructure() []byte {
	var output string
	output += Utils.MakeTitle(d.name, 1)
	output += "<table><tbody>"
	output += Utils.MakeTableRow(true, "Name", "Last modified", "Size")
	output += "<tr><th colspan=\"3\"><hr></th></tr>"
	prePath := d.name[2:]
	if prePath == "/" {
		prePath = ""
	} else {
		output += "<tr><td><a href=\"" + prePath[:maxInt(1, strings.LastIndex(prePath, "/"))] + "\">..</a></td><td></td><td></td></tr>"
	}
	for _, file := range d.fileList {
		size := strconv.Itoa(int(file.Size()))
		if size == "0" {
			size = "-"
		}
		output += Utils.MakeTableRow(false, "<a href=\"" + prePath + "/" + file.Name() + "\">" + file.Name() + "</a>", file.ModTime().Format(time.UnixDate), size)
	}
	output += "</tbody></table>"
	output += "<style>td { padding-right: 10px; }</style>"
	return []byte(output)
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
