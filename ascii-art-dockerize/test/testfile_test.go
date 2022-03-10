package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
	"web"
)

func TestGET(t *testing.T) {
	t.Run("Testing Get request", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		web.GetHandler(response, request)

		tpl, _ := template.ParseFiles("../templates/form.html")

		var bytes bytes.Buffer

		tpl.Execute(&bytes, nil)

		got := response.Body.String()
		want := bytes.String()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
