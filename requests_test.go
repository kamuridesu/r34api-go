package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendGETRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	got, err := sendGETRequest(ts.URL)
	if err != nil {
		t.Fatalf("sendGETRequest failed with error: %v", err)
	}

	want := "Hello, client\n"
	if got != want {
		t.Errorf("sendGETRequest() = %q, want %q", got, want)
	}
}
