package gotumblr

type TextPost struct {
	BasePost
	Title string `json:"title"`
	Body  string `json:"body"`
}

type AltSize struct {
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
	URL    string `json:"url"`
}

type PhotoObject struct {
	Caption  string    `json:"caption"`
	AltSizes []AltSize `json:"alt_sizes"`
}

type PhotoPost struct {
	BasePost
	Photos  []PhotoObject `json:"photos"`
	Caption string        `json:"caption"`
	Width   int64         `json:"width"`
	Height  int64         `json:"height"`
}

type QuotePost struct {
	BasePost
	Text   string `json:"text"`
	Source string `json:"source"`
}

type LinkPost struct {
	BasePost
	Title       string `json:"title"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

type DialogueInfo struct {
	Name   string `json:"name"`
	Label  string `json:"label"`
	Phrase string `json:"phrase"`
}

type ChatPost struct {
	BasePost
	Title    string         `json:"title"`
	Body     string         `json:"body"`
	Dialogue []DialogueInfo `json:"dialogue"`
}

type AudioPost struct {
	BasePost
	Caption     string `json:"caption"`
	Player      string `json:"player"`
	Plays       int64  `json:"plays"`
	AlbumArt    string `json:"album_art"`
	Artist      string `json:"artist"`
	Album       string `json:"album"`
	TrackName   string `json:"track_name"`
	TrackNumber int64  `json:"track_number"`
	Year        int64  `json:"year"`
}

type PlayerInfo struct {
	Width     int64  `json:"width"`
	EmbedCode string `json:"embed_code"`
}

type VideoPost struct {
	BasePost
	Caption string       `json:"caption"`
	Player  []PlayerInfo `json:"player"`
}

type AnswerPost struct {
	BasePost
	AskingName string `json:"asking_name"`
	AskingURL  string `json:"asking_url"`
	Question   string `json:"question"`
	Answer     string `json:"answer"`
}
