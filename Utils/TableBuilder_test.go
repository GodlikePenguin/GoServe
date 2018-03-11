package Utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMakeTableRow_single_row(t *testing.T) {
	singleData := "some-content"
	result := MakeTableRow(false, singleData)
	assert.Equal(t, "<tr><td>" + singleData + "</td></tr>", result)
}

func TestMakeTableRow_multiple_row(t *testing.T) {
	firstData := "some-content"
	secondData := "some-more-content"
	thirdData := "some-final-content"
	result := MakeTableRow(false, firstData, secondData, thirdData)
	assert.Equal(t, "<tr><td>" + firstData + "</td><td>" + secondData + "</td><td>" + thirdData + "</td></tr>", result)
}

func TestMakeTableRow_single_header(t *testing.T) {
	singleData := "some-content"
	result := MakeTableRow(true, singleData)
	assert.Equal(t, "<tr><th>" + singleData + "</th></tr>", result)
}

func TestMakeTableRow_multiple_header(t *testing.T) {
	firstData := "some-content"
	secondData := "some-more-content"
	thirdData := "some-final-content"
	result := MakeTableRow(true, firstData, secondData, thirdData)
	assert.Equal(t, "<tr><th>" + firstData + "</th><th>" + secondData + "</th><th>" + thirdData + "</th></tr>", result)
}

func TestMakeTable(t *testing.T) {
	content := "some-content"
	response := MakeTable(content)
	assert.Equal(t, "<table><tbody>" + content + "</tbody></table>", response)
}
