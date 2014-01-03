package gotumblr

//defines a Go Client for the Tumblr API
type TumblrRestClient struct {
	consumerKey string
	consumerSecret string
	oauthToken string
	oauthSecret string
	host string
	request TumblrRequest
}

//Initializes the TumblrRestClient, creating TumblrRequest that deals with all request formatting.
//consumerKey is the consumer key of your Tumblr Application
//consumerSecret is the consumer secret of your Tumblr Application
//oauthToken is the user specific token, received from the /access_token endpoint
//oauthSecret is the user specific secret, received from the /access_token endpoint
//host is the host that you are tryng to send information to (e.g. http://api.tumblr.com)
func NewTumblrRestClient(consumerKey string, consumerSecret string, oauthToken string, oauthSecret string, host string) *TumblrRestClient {
	
}

//Gets the user information
func (trc TumblrRestClient) Info() map[string]string {

}

//Retrieves the url of the blog's avatar
//size can be: 16, 24, 30, 40, 48, 64, 96, 128 or 512
func (trc TumblrRestClient) Avatar(blogname string, size int) map[string]string {

}

//Gets the likes of the given user
//options can be:
//limit: the number of results to return, inclusive
//offset: liked post number to start at
func (trc TumblrRestClient) Likes(options map[string]int) map[string]string {

}

//Gets the blogs that the user is following.
//options can be:
//limit: the number of results to return
//offset: result number to start at
func (trc TumblrRestClient) Following(options map[string]int) map[string]string {

}

//Gets the dashboard of the user
//options can be:
//limit: number of results to return
//offset: post number to start at
//type: the type of posts to return(text, photo, quote, link, chat, audio, video, answer)
//since_id: return posts that have apeared after this id
//reblog_info: whether to return reblog information about the posts
//notes_info: whether to return notes information about the posts 
func (trc TumblrRestClient) Dashboard(options map[string]string) map[string]string {

}

//Gets a list of posts with the given tag
//tag: the tag you want to look for
//options can be:
//before: the timestamp of when you'd like to see posts before
//limit: the number of results to return
//filter: the post format you want to get(e.g html, text, raw)
func (trc TumblrRestClient) Tagged(tag string, options map[string]string) map[string]string {

}

//Gets a list of posts from a blog
//blogname: the name of the blog you want to get posts from (e.g. mariaterzieva.tumblr.com)
//type: the type of the posts you want to get
//(e.g. text, quote, link, answer, video, audio, photo, chat, all)
//options can be:
//id: the id of the post you are looking for
//tag: return only posts with this tag
//limit: the number of posts to return
//offset: the number of the post you want to start from
//filter: return only posts with a specific format(e.g. html, text, raw)
func (trc TumblrRestClient) Posts(blogname string, type string, options map[string]string) map[string]string {

}

//Gets general information about the blog
//blogname: name of the blog you want to get information about(e.g. mariaterzieva.tumblr.com)
func (trc TumblrRestClient) BlogInfo(blogname string) map[string]string {

}

//Gets the followers of the blog given
//blogname: name of the blog whose followers you want to get
//optons can be:
//limit: the number of results to return, inclusive
//offset: result to start at
func (trc TumblrRestClient) Followers(blogname string, options map[string]int) map[string]string {

}

//Gets the likes of blog given
//blogname: name of the blog whose likes you want to get
//options can be:
//limit: how many likes do you want to get
//offset: the number of the like you want to start from
func (trc TumblrRestClient) BlogLikes(blogname string, options map[string]int) map[string]string {

}

//Gets posts that are currently in the blog's queue
//options can be:
//limit: the number of results to return
//offset: post number to start at
//filter: specify posts' format(e.g. format="html", format="text", format="raw")
func (trc TumblrRestClient) Queue(blogname string, options map[string]string) map[string]string {

}

//Gets posts that are currently in the blog's drafts
//options can be:
//filter: specify posts' format(e.g. format="html", format="text", format="raw")
func (trc TumblrRestClient) Drafts(blogname string, options map[string]string) map[string]string {

}

//Retrieve submission posts
//options can be:
//offset: post number to start at
//filter: specify posts' format(e.g. format="html", format="text", format="raw")
func (trc TumblrRestClient) Submission(blogname string, options map[string]string) map[string]string {

}

//Follow the url of a given blog
//blogname: the url of the blog to follow
func (trc TumblrRestClient) Follow(blogname string) map[string]string {

}

//Unfollow the url of a given blog
//blogname: the url of the blog to unfollow
func (trc TumblrRestClient) Unfollow(blogname string) map[string]string {

}

//Like post of a given blog
//id: the id of the post you want to like
//reblogKey: the reblog key for the post id
func (trc TumblrRestClient) Like(id int, reblogKey string) map[string]string {

}

//Unlike a post of a given blog
//id: the id of the post you want to unlike
//reblogKey: the reblog key for the post id
func (trc TumblrRestClient) Unlike(id int, reblogKey string) map[string]string {

}

//Create a photo post or photoset on a blog
//blogname: the url of the blog you want to post to
//options can be:
//state: the state of the post(e.g. published, draft, queue, private)
//tags: a list of tags you want applied to the post
//tweet: manages the autotweet for this post: set to off for no tweet
//		 or enter text to override the default tweet
//date: the GMT date and time of the post as a string
//format: sets the format type of the post(html or markdown)
//slug: add a short text summary to the end of the post url
//caption: the caption that you want applied to the photo
//link: the 'click-through' url for the photo
//source: the photo source url(either source or data)
//data: one or more image files(wither source or data) 
func (trc TumblrRestClient) CreatePhoto(blogname string, options map[string]string) map[string]string {

}

//Create a text post on a blog
//blogname: the url of the blog you want to post to
//options can be:
//state: the state of the post(e.g. published, draft, queue, private)
//tags: a list of tags you want applied to the post
//tweet: manages the autotweet for this post: set to off for no tweet
//		 or enter text to override the default tweet
//date: the GMT date and time of the post as a string
//format: sets the format type of the post(html or markdown)
//slug: add a short text summary to the end of the post url
//title: the optional title of the post
//body: the full text body
func (trc TumblrRestClient) CreateText(blogname string, options map[string]string) map[string]string {

}

//Create a quote post on a blog
//blogname: the url of the blog you want to post to
//options can be:
//state: the state of the post(e.g. published, draft, queue, private)
//tags: a list of tags you want applied to the post
//tweet: manages the autotweet for this post: set to off for no tweet
//		 or enter text to override the default tweet
//date: the GMT date and time of the post as a string
//format: sets the format type of the post(html or markdown)
//slug: add a short text summary to the end of the post url
//quote: the full text of the quote
//source: the cited source of the quote
func (trc TumblrRestClient) CreateQuote(blogname string, options map[string]string) map[string]string {

}

//Create a link post on a blog
//blogname: the url of the blog you want to post to
//options can be:
//state: the state of the post(e.g. published, draft, queue, private)
//tags: a list of tags you want applied to the post
//tweet: manages the autotweet for this post: set to off for no tweet
//		 or enter text to override the default tweet
//date: the GMT date and time of the post as a string
//format: sets the format type of the post(html or markdown)
//slug: add a short text summary to the end of the post url
//title: the title of the page the link points to
//url: the link you are posting
//description: the description of the link you are posting
func (trc TumblrRestClient) CreateLink(blogname string, options map[string]string) map[string]string {

}

//Create a chat post on a blog
//blogname: the url of the blog you want to post to
//options can be:
//state: the state of the post(e.g. published, draft, queue, private)
//tags: a list of tags you want applied to the post
//tweet: manages the autotweet for this post: set to off for no tweet
//		 or enter text to override the default tweet
//date: the GMT date and time of the post as a string
//format: sets the format type of the post(html or markdown)
//slug: add a short text summary to the end of the post url
//title: the title of the chat
//conversation: the text of the conversation/chat, with dialogue labels
func (trc TumblrRestClient) CreateChatPost(blogname string, options map[string]string) map[string]string {

}

//Create an audio post on a blog
//blogname: the url of the blog you want to post to
//options can be:
//state: the state of the post(e.g. published, draft, queue, private)
//tags: a list of tags you want applied to the post
//tweet: manages the autotweet for this post: set to off for no tweet
//		 or enter text to override the default tweet
//date: the GMT date and time of the post as a string
//format: sets the format type of the post(html or markdown)
//slug: add a short text summary to the end of the post url
//caption: the caption of the post
//external_url: the url of the site that hosts the oudio file
//data: the local filename path to the audio you are uploading
func (trc TumblrRestClient) CreateAudio(blogname string, options map[string]string) map[string]string {

}

//Create a video post on a blog
//blogname: the url of the blog you want to post to
//options can be:
//state: the state of the post(e.g. published, draft, queue, private)
//tags: a list of tags you want applied to the post
//tweet: manages the autotweet for this post: set to off for no tweet
//		 or enter text to override the default tweet
//date: the GMT date and time of the post as a string
//format: sets the format type of the post(html or markdown)
//slug: add a short text summary to the end of the post url
//caption: the caption for the post
//embed: the html embed code for the video
//data: the local filename path to the video you are uploading
func (trc TumblrRestClient) CreateVideo(blogname string, options map[string]string) map[string]string {

}

//Creates a reblog on the given blog
//blogname: the url of the blog you want to reblog to
//id: the id of the reblogged post
//reblogKey: the reblog key for the post
func (trc TumblrRestClient) Reblog(blogname string, options map[string]string) map[string]string {

}

//Deletes a post with a given id
//blogname: the url of the blog you want to delete from
//id: the id of the post you want to delete
func (trc TumblrRestClient) DeletePost(blogname string, id int) map[string]string {

}

//Edits a post with a given id
//blogname: the url of the blog you want to post to
//options can be:
//tags: a list of tags you want applied to the post
//tweet: manages the autotweet for this post: set to off for no tweet
//		 or enter text to override the default tweet
//date: the GMT date and time of the post as a string
//format: sets the format type of the post(html or markdown)
//slug: add a short text summary to the end of the post url
func (trc TumblrRestClient) EditPost(blogname string, options map[string]string) map[string]string {

}

//Formats parameters and sends the API requests. Validates common parameters and formats your tags for you.
//blogname: the name of the blog you are posting to
//params: map of the parameters for the API request
//valid_options: list of valid options that the request allows
func (trc TumblrRestClient) sendPost(blogname string, params map[string]string, validOptions []string) map[string]string {

}

//Sends parameters to the requested url, validating them to make sure that they are what we expect to be passed to us.
//method: the request method you want to make
//params: the parameters used for the API request
//validParams: the list of the valid parameters
//needsApiKey: determines if your request needs an API key 
func (trc TumblrRestClient) sendApiRequest(method string, url string, params map[string]string, validParams []string, needsApiKey bool) map[string]string {

}