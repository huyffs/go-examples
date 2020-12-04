package service

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func successServer(msg string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, msg)
	}))
}

func errorServer(code int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
	}))
}

func TestCallRemoteEndpointOK(t *testing.T) {
	ts := successServer("Hello, client")
	defer ts.Close()

	_, err := CallRemoteEndpoint(ts.URL)
	if err != nil {
		t.Errorf("Expected success, got error: %v", err)
	}
}

func TestCallRemoteEndpoint400(t *testing.T) {
	ts := errorServer(http.StatusBadRequest)
	defer ts.Close()

	t.Log("Testing 400")
	data, err := CallRemoteEndpoint(ts.URL)
	if err == nil {
		t.Errorf("Expected error, got success: data=%v", data)
	}
}

func TestCallRemoteEndpoint404(t *testing.T) {
	ts := errorServer(http.StatusNotFound)
	defer ts.Close()

	t.Log("Testing 404")
	data, err := CallRemoteEndpoint(ts.URL)
	if err == nil {
		t.Errorf("Expected error, got success: data=%v", data)
	}
}
