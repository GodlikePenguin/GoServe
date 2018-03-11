package Utils

func MakeTable(contents string) string {
	return "<table><tbody>" + contents + "</tbody></table>"
}

func MakeTableRow(header bool, items ...string) string {
	var dataTypeStart string
	var dataTypeEnd string
	if header {
		dataTypeStart = "<th>"
		dataTypeEnd = "</th>"
	} else {
		dataTypeStart = "<td>"
		dataTypeEnd = "</td>"
	}

	output := "<tr>"
	for _, data := range items {
		output += dataTypeStart + data + dataTypeEnd
	}
	output += "</tr>"
	return output
}
