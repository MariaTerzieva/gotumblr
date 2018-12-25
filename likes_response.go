package gotumblr

import (
	"encoding/json"
)

type LikesResponse struct {
	LikedPosts []json.RawMessage `json:"liked_posts"`
	LikedCount int64             `json:"liked_count"`
}
