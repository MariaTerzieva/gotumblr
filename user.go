package gotumblr

type User struct {
	Name      string `json:"name"`
	Following bool   `json:"following"`
	URL       string `json:"url"`
	Updated   int64  `json:"updated"`
}
