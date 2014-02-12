package gotumblr

type UserInfo struct {
	Following           int64
	Default_post_format string
	Name                string
	Likes               int64
	Blogs               []OwnedBlog
}
