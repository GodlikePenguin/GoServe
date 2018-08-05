package DirHandler

import (
	"github.com/GodlikePenguin/GoServe/Utils"
	"github.com/davecgh/go-spew/spew"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const mytemplate = `<!DOCTYPE html>
<html>
<head>
    <title>Go Serve</title>
    <meta charset="UTF-8">
    <style>td { padding-right: 10px; }</style>
</head>
<body>
<h1>{{.Name}}</h1>
<table>
    <tbody>
        <tr>
            <th>Name</th>
            <th>Last Modified</th>
            <th>Size</th>
        </tr>
        <tr><th colspan="3"><hr></th></tr>
        <tr><td><a href="{{.BackPath}}">..</a></td><td></td><td></td></tr>
    {{range .Files}}
        <tr>
            <td><a href="{{.Location}}">{{.Name}}</a></td>
            <td>{{.LastModified}}</td>
            <td>{{.Size}}</td>
        </tr>
    {{end}}
    </tbody>
</table>
</body>
</html>`

const mytemplatenoback = `<!DOCTYPE html>
<html>
<head>
    <title>Go Serve</title>
    <meta charset="UTF-8">
    <style>td { padding-right: 10px; }</style>
</head>
<body>
<h1>{{.Name}}</h1>
<table>
    <tbody>
        <tr>
            <th>Name</th>
            <th>Last Modified</th>
            <th>Size</th>
        </tr>
        <tr><th colspan="3"><hr></th></tr>
    {{range .Files}}
        <tr>
            <td><a href="{{.Location}}">{{.Name}}</a></td>
            <td>{{.LastModified}}</td>
            <td>{{.Size}}</td>
        </tr>
    {{end}}
    </tbody>
</table>
</body>
</html>`

func DirHandler(w http.ResponseWriter, r *http.Request, directory string, path string) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Printf(err.Error())
		Utils.ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	spew.Dump(path)
	backPath := GetBackPathFor(path)
	spew.Dump(backPath)
	structure := CreateDirectory(directory, backPath, files, path)

	var t *template.Template
	if path != "/" {
		t = template.Must(template.New("html-dir").Parse(mytemplate))
	} else {
		t = template.Must(template.New("html-dir").Parse(mytemplatenoback))
	}
	err = t.Execute(w, structure)
	if err != nil {
		spew.Dump(err)
		Utils.ErrorHandler(w, r, http.StatusInternalServerError)
	}
}

func GetBackPathFor(path string) string {
	index := strings.LastIndex(path, "/")
	if index > 0 {
		return path[:index]
	}
	return "/"
}
