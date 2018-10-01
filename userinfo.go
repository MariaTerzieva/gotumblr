package gotumblr

type UserInfo struct {
	Following         int64       `json:"following"`
	DefaultPostFormat string      `json:"default_post_format"`
	Name              string      `json:"name"`
	Likes             int64       `json:"likes"`
	Blogs             []OwnedBlog `json:"blogs"`
}
