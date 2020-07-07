package demoapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"os"
)

var a App

func TestMain(m *testing.M) {
	a = App{}
	a.Initialize()

	code := m.Run()

	os.Exit(code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func TestGetHello(t *testing.T) {	
	req, _ := http.NewRequest("GET", "/hello/Leozinho", nil)
	response := executeRequest(req)

	if body := response.Body.String(); body != "Ola Leozinho" {
		t.Errorf("Expected Ola Leozinho, got %s", body)
	}
}

func TestGetSum(t *testing.T) {
	req, _ := http.NewRequest("GET", "/sum/1/2", nil)
	response := executeRequest(req)

	if body := response.Body.String(); body != "3" {
		t.Errorf("Expected 3, got %s", body)
	}
}
