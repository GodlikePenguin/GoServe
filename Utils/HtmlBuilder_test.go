package Utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strconv"
)

func TestMakeTitle(t *testing.T) {
	contents := "my Title"
	level := 1
	result := MakeTitle(contents, level)
	assert.Equal(t, "<h" + strconv.Itoa(level) + ">" + contents + "</h" + strconv.Itoa(level) + ">", result)
}
