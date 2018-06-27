package gotumblr

import "encoding/json"

type NotesResponse struct {
	Notes []json.RawMessage
}
