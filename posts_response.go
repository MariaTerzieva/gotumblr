package gotumblr

import "encoding/json"

type PostsResponse struct {
	Blog  BlogInfo
	Posts []json.RawMessage
	Total_posts int64
}
