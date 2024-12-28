package main

import (
	"fmt"
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
