package linkedin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// LinkedInClient is the main API client.
type LinkedInClient struct {
	HTTPClient  *http.Client
	AccessToken string
	BaseURL     string
}

// NewClient initializes a new LinkedIn API client.
func NewClient(accessToken string) *LinkedInClient {
	return &LinkedInClient{
		HTTPClient:  &http.Client{},
		AccessToken: accessToken,
		BaseURL:     "https://api.linkedin.com/v2",
	}
}

// Get makes a GET request to the LinkedIn API.
func (c *LinkedInClient) Get(endpoint string, result interface{}) error {
	req, err := http.NewRequest("GET", c.BaseURL+endpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("LinkedIn API error: %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(result)
}

// Post makes a POST request to the LinkedIn API.
func (c *LinkedInClient) Post(endpoint string, data interface{}) error {
	body, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", c.BaseURL+endpoint, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("LinkedIn API error: %d", resp.StatusCode)
	}
	return nil
}
