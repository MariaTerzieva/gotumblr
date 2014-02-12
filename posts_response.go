package gotumblr

import "encoding/json"

type PostsResponse struct {
	Blog  BlogInfo
	Posts []json.RawMessage
}
