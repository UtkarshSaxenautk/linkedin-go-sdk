package linkedin

type PostContent struct {
	Author         string `json:"author"`
	LifecycleState string `json:"lifecycleState"`
	Visibility     string `json:"visibility"`
	Text           struct {
		Text string `json:"text"`
	} `json:"specificContent"`
}

// CreatePost posts an update on LinkedIn.
func (c *LinkedInClient) CreatePost(text string) error {
	data := PostContent{
		Author:         "urn:li:person:yourLinkedInID",
		LifecycleState: "PUBLISHED",
		Visibility:     "PUBLIC",
	}
	data.Text.Text = text

	return c.Post("/ugcPosts", data)
}
