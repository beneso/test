package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)



func TestHealthHandler(t *testing.T) {
	expected := "K"
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()
	healthHandler(w, req)
	res := w.Result()
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if string(data) != expected {
		t.Errorf("Expected OK but got %v", string(data))
	}
}
