package gotumblr

type FollowedBlog struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	Updated     int64  `json:"updated"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
