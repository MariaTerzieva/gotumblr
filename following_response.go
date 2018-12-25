package gotumblr

type FollowingResponse struct {
	TotalBlogs int64          `json:"total_blogs"`
	Blogs      []FollowedBlog `json:"blogs"`
}
