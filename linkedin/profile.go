package linkedin

// Profile represents a LinkedIn user profile.
type Profile struct {
	ID        string `json:"id"`
	FirstName string `json:"localizedFirstName"`
	LastName  string `json:"localizedLastName"`
}

// GetProfile fetches the authenticated user's LinkedIn profile.
func (c *LinkedInClient) GetProfile() (*Profile, error) {
	var profile Profile
	err := c.Get("/me", &profile)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}
