package linkedin

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/linkedin"
)

// LinkedInAuth handles OAuth2 authentication for LinkedIn.
type LinkedInAuth struct {
	Config *oauth2.Config
}

// NewLinkedInAuth initializes a new LinkedInAuth with the provided credentials.
func NewLinkedInAuth(clientID, clientSecret, redirectURL string) *LinkedInAuth {
	return &LinkedInAuth{
		Config: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  redirectURL,
			Scopes:       []string{"r_liteprofile", "r_emailaddress", "w_member_social"},
			Endpoint:     linkedin.Endpoint,
		},
	}
}

// GetAuthURL generates the LinkedIn login URL with a secure state token.
// It returns the URL and the generated state string.
func (auth *LinkedInAuth) GetAuthURL() (string, string) {
	state := generateState()
	return auth.Config.AuthCodeURL(state, oauth2.AccessTypeOffline), state
}

// ExchangeCode exchanges an authorization code for an access token.
func (auth *LinkedInAuth) ExchangeCode(ctx context.Context, code string) (*oauth2.Token, error) {
	if code == "" {
		return nil, errors.New("authorization code is empty")
	}
	token, err := auth.Config.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange auth code: %w", err)
	}
	if token.AccessToken == "" {
		return nil, errors.New("received empty access token")
	}
	return token, nil
}

// HandleCallback is a helper HTTP handler to process the OAuth2 callback.
// It exchanges the code for a token and then fetches the user profile.
func (auth *LinkedInAuth) HandleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Missing authorization code", http.StatusBadRequest)
		return
	}

	token, err := auth.ExchangeCode(r.Context(), code)
	if err != nil {
		http.Error(w, "Failed to exchange code: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a client with the access token and fetch profile info.
	client := NewClient(token.AccessToken)
	profile, err := client.GetProfile()
	if err != nil {
		http.Error(w, "Failed to fetch profile: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := jsonResponse(w, profile); err != nil {
		http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

// generateState creates a random state string for security.
func generateState() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

// jsonResponse writes the given data as JSON to the ResponseWriter.
func jsonResponse(w http.ResponseWriter, data interface{}) error {
	return json.NewEncoder(w).Encode(data)
}
