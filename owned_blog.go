package gotumblr

type OwnedBlog struct {
	Name      string
	Url       string
	Title     string
	Primary   bool
	Followers int64
	Tweet     string
	Facebook  string
	Type      string
}
