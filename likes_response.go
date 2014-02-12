package gotumblr

import "encoding/json"

type LikesResponse struct {
	Liked_posts []json.RawMessage
	Liked_count int64
}
