package gotumblr

import "encoding/json"

type BasePost struct {
	BlogName                string        `json:"blog_name"`
	ID                      json.Number   `json:"id"`
	PostURL                 string        `json:"post_url"`
	PostType                string        `json:"type"`
	Timestamp               int64         `json:"timestamp"`
	Date                    string        `json:"date"`
	Format                  string        `json:"format"`
	ReblogKey               string        `json:"reblog_key"`
	Tags                    []string      `json:"tags"`
	Bookmarklet             bool          `json:"bookmarklet"`
	Mobile                  bool          `json:"mobile"`
	SourceURL               string        `json:"source_url"`
	SourceTitle             string        `json:"source_title"`
	Liked                   bool          `json:"liked"`
	State                   string        `json:"state"`
	TotalPosts              int64         `json:"total_posts"`
	NoteCount               int64         `json:"note_count"`
	Notes                   []Note        `json:"notes"`
	Reblog                  ReblogComment `json:"reblog"`
	RebloggedFromID         json.Number   `json:"reblogged_from_id"`
	RebloggedFromURL        string        `json:"reblogged_from_url"`
	RebloggedFromName       string        `json:"reblogged_from_name"`
	RebloggedFromTitle      string        `json:"reblogged_from_title"`
	RebloggedFromUUID       string        `json:"reblogged_from_uuid"`
	RebloggedFromCanMessage bool          `json:"reblogged_from_can_message"`
	RebloggedRootID         json.Number   `json:"reblogged_root_id"`
	RebloggedRootURL        string        `json:"reblogged_root_url"`
	RebloggedRootName       string        `json:"reblogged_root_name"`
	RebloggedRootTitle      string        `json:"reblogged_root_title"`
	RebloggedRootUUID       string        `json:"reblogged_root_uuid"`
	RebloggedRootCanMessage bool          `json:"reblogged_root_can_message"`
}

type ReblogComment struct {
	Comment  string `json:"comment"`
	TreeHTML string `json:"tree_html"`
}
