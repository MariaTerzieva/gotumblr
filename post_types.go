package gotumblr

type TextPost struct {
	BasePost
	Title, Body string
}

type AltSize struct {
	Width, Height int64
	Url           string
}

type PhotoObject struct {
	Caption   string
	Alt_sizes []AltSize
}

type PhotoPost struct {
	BasePost
	Photos        []PhotoObject
	Caption       string
	Width, Height int64
}

type QuotePost struct {
	BasePost
	Text, Source string
}

type LinkPost struct {
	BasePost
	Title, Url, Description string
}

type DialogueInfo struct {
	Name, Label, Phrase string
}

type ChatPost struct {
	BasePost
	Title, Body string
	Dialogue    []DialogueInfo
}

type AudioPost struct {
	BasePost
	Caption      string
	Player       string
	Plays        int64
	Album_art    string
	Artist       string
	Album        string
	Track_name   string
	Track_number int64
	Year         int64
}

type PlayerInfo struct {
	Width      int64
	Embed_code string
}

type VideoPost struct {
	BasePost
	Caption string
	Player  []PlayerInfo
}

type AnswerPost struct {
	BasePost
	Asking_name string
	Asking_url  string
	Question    string
	Answer      string
}
