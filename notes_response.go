package gotumblr

import "encoding/json"

type NotesResponse struct {
	Notes []Note `json:"notes"`
}

type Note struct {
	Type                 string      `json:"type"`
	Timestamp            int64       `json:"timestamp"`
	BlogName             string      `json:"blog_name"`
	BlogUUID             string      `json:"blog_uuid"`
	BlogURL              string      `json:"blog_url"`
	Followed             bool        `json:"followed"`
	AvatarShape          string      `json:"avatar_shape"`
	ReplyText            string      `json:"reply_text"`
	PostID               json.Number `json:"post_id"`
	ReblogParentBlogName string      `json:"reblog_parent_blog_name"`
}
