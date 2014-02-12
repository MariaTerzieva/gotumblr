package gotumblr

type FollowingResponse struct {
	Total_blogs int64
	Blogs       []FollowedBlog
}
