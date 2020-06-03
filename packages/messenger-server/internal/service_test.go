package internal

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestHelloWorldHandler(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	helloWorldHandler(rr, req)

	if rr.Result().StatusCode != 200 {
		t.Fatalf("Expected 200, got %d", rr.Result().StatusCode)
	}

	respBody, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	body := string(respBody)

	if body != "Hello world" {
		t.Fatalf("Expected %s, got %s", "Hello world", body)
	}
}

func TestHelloAnyHandler(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	req = mux.SetURLVars(req, map[string]string{"name": "Me"})

	helloAnyHandler(rr, req)

	if rr.Result().StatusCode != 200 {
		t.Fatalf("Expected 200, got %d", rr.Result().StatusCode)
	}

	respBody, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	body := string(respBody)

	if body != "Hello Me" {
		t.Fatalf("Expected %s, got %s", "Hello Me", body)
	}
}
