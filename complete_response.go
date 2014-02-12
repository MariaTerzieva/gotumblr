package gotumblr

import "encoding/json"

type CompleteResponse struct {
	Meta     MetaInfo
	Response json.RawMessage
}
