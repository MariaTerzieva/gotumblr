package gotumblr

type BasePost struct {
	Blog_name    string
	Id           int64
	Post_url     string
	PostType     string `json:"type"`
	Timestamp    int64
	Date         string
	Format       string
	Reblog_key   string
	Tags         []string
	Bookmarklet  bool
	Mobile       bool
	Source_url   string
	Source_title string
	Liked        bool
	State        string
	Total_Posts  int64
	Note_count   int64
	Notes        []Note
	Photos       []Photo
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

type Photo struct {
	Caption                 string
	Alt_sizes               []Img_size
	Original_size           Img_size
}

type Img_size struct {
	Width                	int64
	Height              	int64
	Url						string
}
