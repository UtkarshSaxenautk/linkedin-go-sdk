package linkedin_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/UtkarshSaxenautk/linkedin-go-sdk/linkedin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestNewLinkedInAuth(t *testing.T) {
	auth := linkedin.NewLinkedInAuth("client_id", "client_secret", "http://localhost/callback")
	require.NotNil(t, auth)
	assert.Equal(t, "client_id", auth.Config.ClientID)
	assert.Equal(t, "client_secret", auth.Config.ClientSecret)
	assert.Equal(t, "http://localhost/callback", auth.Config.RedirectURL)
}

func TestGetAuthURL(t *testing.T) {
	auth := linkedin.NewLinkedInAuth("client_id", "client_secret", "http://localhost/callback")
	url, state := auth.GetAuthURL()

	require.NotEmpty(t, url)
	require.NotEmpty(t, state)
	assert.Contains(t, url, "https://www.linkedin.com/oauth")
}

func TestExchangeCode(t *testing.T) {
	auth := linkedin.NewLinkedInAuth("client_id", "client_secret", "http://localhost/callback")
	ctx := context.Background()

	// Simulate LinkedIn's token exchange endpoint
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"access_token": "test_token",
		})
	}))
	defer server.Close()
	auth.Config.Endpoint.TokenURL = server.URL
	token, err := auth.ExchangeCode(ctx, "test_code")
	require.NoError(t, err)
	require.NotNil(t, token)
	assert.Equal(t, "test_token", token.AccessToken)
}

// MockLinkedInClient mocks LinkedIn API interactions.
type MockLinkedInClient struct {
	mock.Mock
}

func (m *MockLinkedInClient) GetProfile() (map[string]interface{}, error) {
	args := m.Called()
	return args.Get(0).(map[string]interface{}), args.Error(1)
}
