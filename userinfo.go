package gotumblr

type userInfo struct {
	Following int64
	Default_post_format string
	Name string
	Likes int64
	Blogs []ownedBlog
}