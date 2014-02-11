package gotumblr

import "encoding/json"

type completeResponse struct {
	Meta meta
	Response json.RawMessage
}