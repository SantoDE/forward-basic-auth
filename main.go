package main

import (
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", http.HandlerFunc(basicAuth)); err != nil {
		log.Fatal(err)
	}
}

func basicAuth(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if ok {
		usernameHash := sha256.Sum256([]byte(username))
		passwordHash := sha256.Sum256([]byte(password))
		expectedUsernameHash := sha256.Sum256([]byte("test"))
		expectedPasswordHash := sha256.Sum256([]byte("test1"))

		usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
		passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

		if usernameMatch && passwordMatch {
			w.Write([]byte("Passed"))
			w.WriteHeader(200)
			return
		}
	}

	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}