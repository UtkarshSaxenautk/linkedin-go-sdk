package linkedin

import (
	"context"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/linkedin"
)

// LinkedInAuth handles OAuth2 authentication.
type LinkedInAuth struct {
	Config *oauth2.Config
}

// NewLinkedInAuth initializes LinkedIn OAuth2 configuration.
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

// GetAuthURL returns the authorization URL for user login.
func (auth *LinkedInAuth) GetAuthURL(state string) string {
	return auth.Config.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

// ExchangeCode exchanges an authorization code for an access token.
func (auth *LinkedInAuth) ExchangeCode(ctx context.Context, code string) (*oauth2.Token, error) {
	return auth.Config.Exchange(ctx, code)
}
