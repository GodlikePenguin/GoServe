package main

import (
	"testing"
	"net/http"
	"os"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
)

func TestCustomFileServer_valid_dir(t *testing.T) {
	originalValue := os.Getenv("SERVE_LOCATION")
	os.Setenv("SERVE_LOCATION", "D:/")
	os.MkdirAll("D:/tmp/dir1", os.FileMode(0522))
	os.MkdirAll("D:/tmp/dir2", os.FileMode(0522))
	defer os.RemoveAll("D:/tmp")
	defer os.Setenv("SERVE_LOCATION", originalValue)
	req, err := http.NewRequest("GET", "/tmp", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CustomFileServer)

	handler.ServeHTTP(rr, req)

	status := rr.Code
	assert.Equal(t, 200, status)
	assert.Contains(t, rr.Body.String(), "dir1")
	assert.Contains(t, rr.Body.String(), "dir2")
	assert.NotContains(t, rr.Body.String(), "dir3")
}

func TestCustomFileServer_invalid_dir(t *testing.T) {
	req, err := http.NewRequest("GET", "/invalid-dir", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CustomFileServer)

	handler.ServeHTTP(rr, req)

	status := rr.Code
	assert.Equal(t, 404, status)
	assert.Equal(t, "<h1>404 File Not Found</h1>", rr.Body.String())
}
