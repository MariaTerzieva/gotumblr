package gotumblr

import "encoding/json"

type postsResponse struct {
	Blog blog
	Posts []json.RawMessage
}