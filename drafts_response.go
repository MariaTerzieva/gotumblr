package gotumblr

import "encoding/json"

type DraftsResponse struct {
	Posts []json.RawMessage
}
