package gotumblr

type BasePost struct {
	Blog_name                  string
	Id                         int64
	Post_url                   string
	PostType                   string `json:"type"`
	Timestamp                  int64
	Date                       string
	Format                     string
	Reblog_key                 string
	Tags                       []string
	Bookmarklet                bool
	Mobile                     bool
	Source_url                 string
	Source_title               string
	Liked                      bool
	State                      string
	Total_Posts                int64
	Note_count                 int64
	Notes                      []Note
	Reblogged_from_id          string
	Reblogged_from_url         string
	Reblogged_from_name        string
	Reblogged_from_title       string
	Reblogged_from_uuid        string
	Reblogged_from_can_message bool
	Reblogged_root_id          string
	Reblogged_root_url         string
	Reblogged_root_name        string
	Reblogged_root_title       string
	Reblogged_root_uuid        string
	Reblogged_root_can_message bool
}

type Note struct {
	Type                    string
	Timestamp               int64
	Blog_name               string
	Blog_uuid               string
	Blog_url                string
	Followed                bool
	Avatar_shape            string
	Post_id                 string
	Reblog_parent_blog_name string
}
