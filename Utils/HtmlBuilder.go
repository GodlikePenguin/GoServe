package Utils

import "strconv"

func MakeTitle(contents string, level int) string {
	return "<h" + strconv.Itoa(level) + ">" + contents + "</h" + strconv.Itoa(level) + ">"
}