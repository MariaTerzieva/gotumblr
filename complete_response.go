package gotumblr

import (
	"encoding/json"
)

type CompleteResponse struct {
	Meta     MetaInfo        `json:"meta"`
	Response json.RawMessage `json:"response"`
}
