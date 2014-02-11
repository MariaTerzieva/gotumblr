package gotumblr

import "encoding/json"

type draftsResponse struct {
	Posts []json.RawMessage
}