package gotumblr

import "encoding/json"

type likesResponse struct {
	Liked_posts []json.RawMessage
	Liked_count int64
}