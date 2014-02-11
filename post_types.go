package gotumblr

type textPost struct {
	basePost
	Title, Body string
}

type altSize struct {
	Width, Height int64
	Url string
}

type photoObject struct {
	Caption string
	Alt_sizes []altSize
}

type photoPost struct {
	basePost
	Photos []photoObject
	Caption string
	Width, Height int64
}

type quotePost struct {
	basePost
	Text, Source string
}

type linkPost struct {
	basePost
	Title, Url, Description string
}

type dialogue struct {
	Name, Label, Phrase string
}

type chatPost struct {
	basePost
	Title, Body string
	Dialogue []dialogue
}

type audioPost struct {
	basePost
	Caption string
	Player string
	Plays int64
	Album_art string
	Artist string
	Album string
	Track_name string
	Track_number int64
	Year int64
}

type player struct {
	Width int64
	Embed_code string
}

type videoPost struct {
	basePost
	Caption string
	Player []player
}

type answerPost struct {
	basePost
	Asking_name string
	Asking_url string
	Question string
	Answer string
}