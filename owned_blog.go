package gotumblr

type OwnedBlog struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Primary     bool   `json:"primary"`
	Followers   int64  `json:"followers"`
	Tweet       string `json:"tweet"`
	Facebook    string `json:"facebook"`
	Type        string `json:"type"`
}
