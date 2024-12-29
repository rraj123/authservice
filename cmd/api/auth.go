package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

const alphanumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func (app *application) getAuthCode(w http.ResponseWriter, r *http.Request) {
	// Initialize random with a seed
	rand.Seed(time.Now().UnixNano())

	// Generate a 10-character random string
	b := make([]byte, 10)
	for i := range b {
		b[i] = alphanumeric[rand.Intn(len(alphanumeric))]
	}

	fmt.Fprintf(w, "Generated auth code: %s\n", string(b))
}

func (app *application) getRequestInfo(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	defer r.Body.Close()
	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = io.ReadAll(r.Body)
	}

	// Start writing response
	w.Header().Set("Content-Type", "text/plain")

	// Write headers to response
	fmt.Fprintln(w, "Request Headers:")
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Fprintf(w, "%s: %s\n", name, value)
		}
	}

	// Write body to response
	fmt.Fprintln(w, "\nRequest Body:")
	if len(bodyBytes) > 0 {
		fmt.Fprintf(w, "%s\n", string(bodyBytes))
	} else {
		fmt.Fprintln(w, "<empty body>")
	}
}
