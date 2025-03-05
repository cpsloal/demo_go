package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestVersionHandler tests the version handler.
func TestVersionHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/version", nil)
	w := httptest.NewRecorder()
	version(w, req)

	res := w.Result()
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Error reading body: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v, body:%v", res.Status, string(body))

		t.Errorf("Expected status OK; got %v", res.Status)
	}

}

// TestGreetHandler tests the greet handler.
func TestGreetHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/?name=test", nil)
	w := httptest.NewRecorder()
	greet(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %v", res.Status)
	}

}
func TestGreetHandlerEmpty(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	greet(w, req)

	res := w.Result()
	defer res.Body.Close()
}
