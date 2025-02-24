package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/UtkarshSaxenautk/linkedin-go-sdk/linkedin"
)

var auth *linkedin.LinkedInAuth

func main() {
	clientID := os.Getenv("LINKEDIN_CLIENT_ID")
	clientSecret := os.Getenv("LINKEDIN_CLIENT_SECRET")
	redirectURL := "http://localhost:8080/callback"

	if clientID == "" || clientSecret == "" {
		log.Fatal("Please set LINKEDIN_CLIENT_ID and LINKEDIN_CLIENT_SECRET environment variables.")
	}

	auth = linkedin.NewLinkedInAuth(clientID, clientSecret, redirectURL)

	http.HandleFunc("/", loginHandler)
	http.HandleFunc("/callback", callbackHandler)

	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// loginHandler redirects the user to LinkedIn's OAuth page.
func loginHandler(w http.ResponseWriter, r *http.Request) {
	authURL, state := auth.GetAuthURL()
	// In production, store the generated state in a secure session for validation.
	fmt.Printf("Generated state: %s\n", state)
	http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
}

// callbackHandler processes the OAuth callback.
func callbackHandler(w http.ResponseWriter, r *http.Request) {
	auth.HandleCallback(w, r)
}
