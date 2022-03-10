package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"web"
)

type AsciiArtPost struct {
	Ascii_Art_Text string
}

var mux *http.ServeMux
var writer *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	mux = http.NewServeMux()
	mux.HandleFunc("/ascii-art", web.PostHandler)
	writer = httptest.NewRecorder()
}

func TestHandlerGET(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/ascii-art", web.PostHandler)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("Get", "/", nil)

	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	var post AsciiArtPost
	json.Unmarshal(writer.Body.Bytes(), &post)

	if AsciiArtPost.Ascii_Art_Text !=  {
		t.Error("Cannot retrieve JSON post")
	}
}
