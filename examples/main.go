package main

import (
	"context"
	"fmt"
	"log"

	"github.com/UtkarshSaxenautk/linkedin-go-sdk/linkedin"
)

func main() {
	clientID := "YOUR_CLIENT_ID"
	clientSecret := "YOUR_CLIENT_SECRET"
	redirectURL := "http://localhost:8080/callback"

	auth := linkedin.NewLinkedInAuth(clientID, clientSecret, redirectURL)
	fmt.Println("Visit this URL to authenticate:", auth.GetAuthURL("state123"))

	// Assume we received `authCode` from LinkedIn after user login
	authCode := "RECEIVED_AUTH_CODE"
	token, err := auth.ExchangeCode(context.Background(), authCode)
	if err != nil {
		log.Fatal(err)
	}

	client := linkedin.NewClient(token.AccessToken)

	// Fetch Profile
	profile, err := client.GetProfile()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Welcome, %s %s!\n", profile.FirstName, profile.LastName)

	// Post an update
	err = client.CreatePost("Hello LinkedIn from Go!")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Post successful!")
}
