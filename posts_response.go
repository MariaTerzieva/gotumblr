package gotumblr

import (
	"encoding/json"
)

type PostsResponse struct {
	Blog       BlogInfo          `json:"blog"`
	Posts      []json.RawMessage `json:"posts"`
	TotalPosts int64             `json:"total_posts"`
}
