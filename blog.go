package gotumblr

type BlogInfo struct {
	Title       string `json:"title"`
	Posts       int64  `json:"posts"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Updated     int64  `json:"updated"`
	Description string `json:"description"`
	Ask         bool   `json:"ask"`
	AskAnon     bool   `json:"ask_anon"`
	Likes       int64  `json:"likes"`
}
