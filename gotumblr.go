// A Go Tumblr API v2 Client.

package gotumblr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/kurrik/oauth1a"
)

// Client defines a Go Client for the Tumblr API.
type Client struct {
	service    *oauth1a.Service
	userConfig *oauth1a.UserConfig
	client     *http.Client
	headers    map[string]string
	host       string
	apiKey     string
}

// Setter funcs are executed at the end of NewClient.
type Setter func(*Client)

// SetHost sets the host URL of Tumblrs API.
func SetHost(h string) Setter {
	return func(c *Client) {
		c.host = h
	}
}

// SetCallbackURL sets the callback URL.
func SetCallbackURL(cb string) Setter {
	return func(c *Client) {
		c.service.CallbackURL = cb
	}
}

// SetClient sets the used HTTP client. This can be used for caching.
func SetClient(hc *http.Client) Setter {
	return func(c *Client) {
		c.client = hc
	}
}

// SetHeaders sets headers that are send with every request.
func SetHeaders(hs map[string]string) Setter {
	return func(c *Client) {
		c.headers = hs
	}
}

// New initializes the Client.
// consumerKey is the consumer key of your Tumblr Application.
// consumerSecret is the consumer secret of your Tumblr Application.
// oauthToken is the user specific token, received from the /access_token endpoint.
// oauthSecret is the user specific secret, received from the /access_token endpoint.
// host is the host that you are tryng to send information to (e.g. https://api.tumblr.com).
func New(consumerKey, consumerSecret, oauthToken, oauthSecret string, setters ...Setter) *Client {
	service := &oauth1a.Service{
		RequestURL:   "https://www.tumblr.com/oauth/request_token",
		AuthorizeURL: "https://www.tumblr.com/oauth/authorize",
		AccessURL:    "https://www.tumblr.com/oauth/access_token",
		ClientConfig: &oauth1a.ClientConfig{
			ConsumerKey:    consumerKey,
			ConsumerSecret: consumerSecret,
			CallbackURL:    "",
		},
		Signer: new(oauth1a.HmacSha1Signer),
	}
	userConfig := oauth1a.NewAuthorizedConfig(oauthToken, oauthSecret)

	client := &Client{
		service:    service,
		userConfig: userConfig,
		client:     &http.Client{},
		host:       "https://api.tumblr.com",
	}
	for _, setter := range setters {
		setter(client)
	}

	return client
}

// Info gets the user information.
func (c *Client) Info() (UserInfoResponse, error) {
	data, err := c.Get("/v2/user/info", nil)
	if err != nil {
		return UserInfoResponse{}, err
	}
	var result UserInfoResponse
	return result, json.Unmarshal(data.Response, &result)
}

// Avatar retrieves the url of the blog's avatar.
// size can be: 16, 24, 30, 40, 48, 64, 96, 128 or 512.
func (c *Client) Avatar(blogname string, size int) (AvatarResponse, error) {
	data, err := c.Get(fmt.Sprintf("/v2/blog/%s/avatar/%d", blogname, size), nil)
	if err != nil {
		return AvatarResponse{}, err
	}
	var res AvatarResponse
	return res, json.Unmarshal(data.Response, &res)
}

// Likes gets the likes of the given user.
// options can be:
// limit: the number of results to return, inclusive;
// offset: liked post number to start at.
func (c *Client) Likes(options url.Values) (LikesResponse, error) {
	data, err := c.Get("/v2/user/likes", options)
	if err != nil {
		return LikesResponse{}, err
	}
	var res LikesResponse
	return res, json.Unmarshal(data.Response, &res)
}

// Following gets the blogs that the user is following.
// options can be:
// limit: the number of results to return;
// offset: result number to start at.
func (c *Client) Following(options url.Values) (FollowingResponse, error) {
	data, err := c.Get("/v2/user/following", options)
	if err != nil {
		return FollowingResponse{}, err
	}
	var res FollowingResponse
	return res, json.Unmarshal(data.Response, &res)
}

// Dashboard gets the dashboard of the user.
// options can be:
// limit: number of results to return;
// offset: post number to start at;
// type: the type of posts to return(text, photo, quote, link, chat, audio, video, answer);
// sinceID: return posts that have apeared after this id;
// reblog_info: whether to return reblog information about the posts;
// notes_info: whether to return notes information about the posts.
func (c *Client) Dashboard(options url.Values) (DraftsResponse, error) {
	data, err := c.Get("/v2/user/dashboard", options)
	if err != nil {
		return DraftsResponse{}, err
	}
	var res DraftsResponse
	return res, json.Unmarshal(data.Response, &res)
}

// Tagged gets a list of posts with the given tag.
// tag: the tag you want to look for.
// options can be:
// before: the timestamp of when you'd like to see posts before;
// limit: the number of results to return;
// filter: the post format you want to get(e.g html, text, raw).
func (c *Client) Tagged(tag string, options url.Values) ([]json.RawMessage, error) {
	options.Set("tag", tag)
	options.Set("api_key", c.apiKey)
	data, err := c.Get("/v2/tagged", options)
	if err != nil {
		return []json.RawMessage{}, err
	}
	res := []json.RawMessage{}
	return res, json.Unmarshal(data.Response, &res)
}

// Posts gets a list of posts from a blog.
// blogname: the name of the blog you want to get posts from (e.g. mgterzieva.tumblr.com).
// postsType: the type of the posts you want to get
// (e.g. text, quote, link, answer, video, audio, photo, chat, all).
// options can be:
// id: the id of the post you are looking for;
// tag: return only posts with this tag;
// limit: the number of posts to return;
// offset: the number of the post you want to start from;
// filter: return only posts with a specific format(e.g. html, text, raw).
func (c *Client) Posts(blogname, postsType string, options url.Values) (PostsResponse, error) {
	reqURL := fmt.Sprintf("/v2/blog/%s/posts", blogname)
	if postsType != "" {
		reqURL += "/" + postsType
	}
	options.Set("api_key", c.apiKey)
	data, err := c.Get(reqURL, options)
	if err != nil {
		return PostsResponse{}, err
	}
	var res PostsResponse
	return res, json.Unmarshal(data.Response, &res)
}

// Notes gets a list of notes on a post.
// options can be:
// mode: likes, reblogs_with_tags or rollup
// before_timestamp: return only notes created before the given timestamp
func (c *Client) Notes(blogname, postID string, options url.Values) (NotesResponse, error) {
	reqURL := fmt.Sprintf("/v2/blog/%s/notes", blogname)
	options.Set("api_key", c.apiKey)
	options.Set("id", postID)
	data, err := c.Get(reqURL, options)
	if err != nil {
		return NotesResponse{}, err
	}
	var res NotesResponse
	return res, json.Unmarshal(data.Response, &res)
}

// BlogInfo gets general information about the blog.
// blogname: name of the blog you want to get information about(e.g. mgterzieva.tumblr.com).
func (c *Client) BlogInfo(blogname string) (BlogInfoResponse, error) {
	reqURL := fmt.Sprintf("/v2/blog/%s/info", blogname)
	options := url.Values{"api_key": []string{c.apiKey}}
	data, err := c.Get(reqURL, options)
	if err != nil {
		return BlogInfoResponse{}, err
	}
	var res BlogInfoResponse
	return res, json.Unmarshal(data.Response, &res)
}

// Followers gets the followers of the blog given.
// blogname: name of the blog whose followers you want to get.
// optons can be:
// limit: the number of results to return, inclusive;
// offset: result to start at.
func (c *Client) Followers(blogname string, options url.Values) (FollowersResponse, error) {
	reqURL := fmt.Sprintf("/v2/blog/%s/followers", blogname)
	data, err := c.Get(reqURL, options)
	if err != nil {
		return FollowersResponse{}, err
	}
	var res FollowersResponse
	return res, json.Unmarshal(data.Response, &res)
}

// BlogLikes gets the likes of blog given.
// blogname: name of the blog whose likes you want to get.
// options can be:
// limit: how many likes do you want to get;
// offset: the number of the like you want to start from.
func (c *Client) BlogLikes(blogname string, options url.Values) (LikesResponse, error) {
	reqURL := fmt.Sprintf("/v2/blog/%s/likes", blogname)
	options.Set("api_key", c.apiKey)
	data, err := c.Get(reqURL, options)
	if err != nil {
		return LikesResponse{}, err
	}
	var res LikesResponse
	return res, json.Unmarshal(data.Response, &res)
}

// Queue gets posts that are currently in the blog's queue.
// options can be:
// limit: the number of results to return;
// offset: post number to start at;
// filter: specify posts' format(e.g. format="html", format="text", format="raw").
func (c *Client) Queue(blogname string, options url.Values) (DraftsResponse, error) {
	reqURL := fmt.Sprintf("/v2/blog/%s/posts/queue", blogname)
	data, err := c.Get(reqURL, options)
	if err != nil {
		return DraftsResponse{}, err
	}
	var res DraftsResponse
	return res, json.Unmarshal(data.Response, &res)
}

// Drafts gets posts that are currently in the blog's drafts.
// options can be:
// filter: specify posts' format(e.g. format="html", format="text", format="raw").
func (c *Client) Drafts(blogname string, options url.Values) (DraftsResponse, error) {
	reqURL := fmt.Sprintf("/v2/blog/%s/posts/draft", blogname)
	data, err := c.Get(reqURL, options)
	if err != nil {
		return DraftsResponse{}, err
	}
	var res DraftsResponse
	return res, json.Unmarshal(data.Response, &res)
}

// Submission retrieve submission posts.
// options can be:
// offset: post number to start at;
// filter: specify posts' format(e.g. format="html", format="text", format="raw").
func (c *Client) Submission(blogname string, options url.Values) (DraftsResponse, error) {
	reqURL := fmt.Sprintf("/v2/blog/%s/posts/submission", blogname)
	data, err := c.Get(reqURL, options)
	if err != nil {
		return DraftsResponse{}, err
	}
	var res DraftsResponse
	return res, json.Unmarshal(data.Response, &res)
}

// Follow the url of a given blog.
// blogname: the url of the blog to follow.
func (c *Client) Follow(blogname string) error {
	reqURL := fmt.Sprintf("/v2/user/follow")
	params := url.Values{"url": []string{blogname}}
	data, err := c.Post(reqURL, params)
	if err != nil {
		return err
	}
	if data.Meta.Status != 200 {
		return fmt.Errorf(data.Meta.Msg)
	}
	return nil
}

// Unfollow the url of a given blog.
// blogname: the url of the blog to unfollow.
func (c *Client) Unfollow(blogname string) error {
	reqURL := fmt.Sprintf("/v2/user/unfollow")
	params := url.Values{"url": []string{blogname}}
	data, err := c.Post(reqURL, params)
	if err != nil {
		return err
	}
	if data.Meta.Status != 200 {
		return fmt.Errorf(data.Meta.Msg)
	}
	return nil
}

// Like post of a given blog.
// id: the id of the post you want to like.
// reblog_key: the reblog key for the post id.
func (c *Client) Like(id, reblogKey string) error {
	reqURL := fmt.Sprintf("/v2/user/like")
	params := url.Values{"id": []string{id}, "reblog_key": []string{reblogKey}}
	data, err := c.Post(reqURL, params)
	if err != nil {
		return err
	}
	if data.Meta.Status != 200 {
		return fmt.Errorf(data.Meta.Msg)
	}
	return nil
}

// Unlike a post of a given blog.
// id: the id of the post you want to unlike.
// reblog_key: the reblog key for the post id.
func (c *Client) Unlike(id, reblogKey string) error {
	reqURL := fmt.Sprintf("/v2/user/unlike")
	params := url.Values{"id": []string{id}, "reblog_key": []string{reblogKey}}
	data, err := c.Post(reqURL, params)
	if err != nil {
		return err
	}
	if data.Meta.Status != 200 {
		return fmt.Errorf(data.Meta.Msg)
	}
	return nil
}

// CreatePhoto creates a photo post or photoset on a blog.
// blogname: the url of the blog you want to post to.
// options can be:
// (with * are marked required options)
// state: the state of the post(e.g. published, draft, queue, private);
// tags: a list of tags you want applied to the post;
// tweet: manages the autotweet for this post: set to off for no tweet
// or enter text to override the default tweet;
// date: the GMT date and time of the post as a string;
// format: sets the format type of the post(html or markdown);
// slug: add a short text summary to the end of the post url;
// caption: the caption that you want applied to the photo;
// link: the 'click-through' url for the photo;
// *source: the photo source url.
func (c *Client) CreatePhoto(blogname string, options url.Values) error {
	reqURL := fmt.Sprintf("/v2/blog/%s/post", blogname)
	options.Set("type", "photo")
	data, err := c.Post(reqURL, options)
	if err != nil {
		return err
	}
	if data.Meta.Status != 201 {
		return fmt.Errorf(data.Meta.Msg)
	}
	return nil
}

// CreateText creates a text post on a blog.
// blogname: the url of the blog you want to post to.
// options can be:
// (with * are marked required options)
// state: the state of the post(e.g. published, draft, queue, private);
// tags: a list of tags you want applied to the post;
// tweet: manages the autotweet for this post: set to off for no tweet
// or enter text to override the default tweet;
// date: the GMT date and time of the post as a string;
// format: sets the format type of the post(html or markdown);
// slug: add a short text summary to the end of the post url;
// title: the optional title of the post;
// *body: the full text body.
func (c *Client) CreateText(blogname string, options url.Values) error {
	reqURL := fmt.Sprintf("/v2/blog/%s/post", blogname)
	options.Set("type", "text")
	data, err := c.Post(reqURL, options)
	if err != nil {
		return err
	}
	if data.Meta.Status != 201 {
		return fmt.Errorf(data.Meta.Msg)
	}
	return nil
}

// CreateQuote creates a quote post on a blog.
// blogname: the url of the blog you want to post to.
// options can be:
// (with * are marked required options)
// state: the state of the post(e.g. published, draft, queue, private);
// tags: a list of tags you want applied to the post;
// tweet: manages the autotweet for this post: set to off for no tweet
// or enter text to override the default tweet;
// date: the GMT date and time of the post as a string;
// format: sets the format type of the post(html or markdown);
// slug: add a short text summary to the end of the post url;
// *quote: the full text of the quote;
// source: the cited source of the quote.
func (c *Client) CreateQuote(blogname string, options url.Values) error {
	reqURL := fmt.Sprintf("/v2/blog/%s/post", blogname)
	options.Set("type", "quote")
	data, err := c.Post(reqURL, options)
	if err != nil {
		return err
	}
	if data.Meta.Status != 201 {
		return fmt.Errorf(data.Meta.Msg)
	}
	return nil
}

// CreateLink creates a link post on a blog.
// blogname: the url of the blog you want to post to.
// options can be:
// (with * are marked required options)
// state: the state of the post(e.g. published, draft, queue, private);
// tags: a list of tags you want applied to the post;
// tweet: manages the autotweet for this post: set to off for no tweet
// or enter text to override the default tweet;
// date: the GMT date and time of the post as a string;
// format: sets the format type of the post(html or markdown);
// slug: add a short text summary to the end of the post url;
// title: the title of the page the link points to;
// *url: the link you are posting;
// description: the description of the link you are posting.
func (c *Client) CreateLink(blogname string, options url.Values) error {
	reqURL := fmt.Sprintf("/v2/blog/%s/post", blogname)
	options.Set("type", "link")
	data, err := c.Post(reqURL, options)
	if err != nil {
		return err
	}
	if data.Meta.Status != 201 {
		return fmt.Errorf(data.Meta.Msg)
	}
	return nil
}

// CreateChatPost creates a chat post on a blog.
// blogname: the url of the blog you want to post to.
// options can be:
// (with * are marked required options)
// state: the state of the post(e.g. published, draft, queue, private);
// tags: a list of tags you want applied to the post;
// tweet: manages the autotweet for this post: set to off for no tweet
// or enter text to override the default tweet;
// date: the GMT date and time of the post as a string;
// format: sets the format type of the post(html or markdown);
// slug: add a short text summary to the end of the post url;
// title: the title of the chat;
// *conversation: the text of the conversation/chat, with dialogue labels.
func (c *Client) CreateChatPost(blogname string, options url.Values) error {
	reqURL := fmt.Sprintf("/v2/blog/%s/post", blogname)
	options.Set("type", "chat")
	data, err := c.Post(reqURL, options)
	if err != nil {
		return err
	}
	if data.Meta.Status != 201 {
		return fmt.Errorf(data.Meta.Msg)
	}
	return nil
}

// CreateAudio creates an audio post on a blog.
// blogname: the url of the blog you want to post to.
// options can be:
// (with * are marked required options)
// state: the state of the post(e.g. published, draft, queue, private);
// tags: a list of tags you want applied to the post;
// tweet: manages the autotweet for this post: set to off for no tweet
// or enter text to override the default tweet;
// date: the GMT date and time of the post as a string;
// format: sets the format type of the post(html or markdown);
// slug: add a short text summary to the end of the post url;
// caption: the caption of the post;
// *externalURL: the url of the site that hosts the audio file.
func (c *Client) CreateAudio(blogname string, options url.Values) error {
	reqURL := fmt.Sprintf("/v2/blog/%s/post", blogname)
	options.Set("type", "audio")
	data, err := c.Post(reqURL, options)
	if err != nil {
		return err
	}
	if data.Meta.Status != 201 {
		return fmt.Errorf(data.Meta.Msg)
	}
	return nil
}

// CreateVideo creates a video post on a blog.
// blogname: the url of the blog you want to post to.
// options can be:
// (with * are marked required options)
// state: the state of the post(e.g. published, draft, queue, private);
// tags: a list of tags you want applied to the post;
// tweet: manages the autotweet for this post: set to off for no tweet
// or enter text to override the default tweet;
// date: the GMT date and time of the post as a string;
// format: sets the format type of the post(html or markdown);
// slug: add a short text summary to the end of the post url;
// caption: the caption for the post;
// *embed: the html embed code for the video.
func (c *Client) CreateVideo(blogname string, options url.Values) error {
	reqURL := fmt.Sprintf("/v2/blog/%s/post", blogname)
	options.Set("type", "video")
	data, err := c.Post(reqURL, options)
	if err != nil {
		return err
	}
	if data.Meta.Status != 201 {
		return fmt.Errorf(data.Meta.Msg)
	}
	return nil
}

// Reblog creates a reblog on the given blog.
// blogname: the url of the blog you want to reblog to.
// options should be:
// (with * are marked required options)
// *id: the id of the reblogged post;
// *reblog_key: the reblog key of the rebloged post.
func (c *Client) Reblog(blogname string, options url.Values) error {
	reqURL := fmt.Sprintf("/v2/blog/%s/post/reblog", blogname)
	data, err := c.Post(reqURL, options)
	if err != nil {
		return err
	}
	if data.Meta.Status != 201 {
		return fmt.Errorf(data.Meta.Msg)
	}
	return nil
}

// DeletePost deletes a post with a given id.
// blogname: the url of the blog you want to delete from.
// id: the id of the post you want to delete.
func (c *Client) DeletePost(blogname, id string) error {
	reqURL := fmt.Sprintf("/v2/blog/%s/post/delete", blogname)
	params := url.Values{"id": []string{id}}
	data, err := c.Post(reqURL, params)
	if err != nil {
		return err
	}
	if data.Meta.Status != 200 {
		return fmt.Errorf(data.Meta.Msg)
	}
	return nil
}

// EditPost edits a post with a given id.
// blogname: the url of the blog you want to post to.
// options can be:
// (with * are marked required options)
// tags: a list of tags you want applied to the post;
// tweet: manages the autotweet for this post: set to off for no tweet
// or enter text to override the default tweet;
// date: the GMT date and time of the post as a string;
// format: sets the format type of the post(html or markdown);
// slug: add a short text summary to the end of the post url;
// *id: the id of the post.
// The other options are specific to the type of post you want to edit.
func (c *Client) EditPost(blogname string, options url.Values) error {
	reqURL := fmt.Sprintf("/v2/blog/%s/post/edit", blogname)
	data, err := c.Post(reqURL, options)
	if err != nil {
		return err
	}
	if data.Meta.Status != 200 {
		return fmt.Errorf(data.Meta.Msg)
	}
	return nil
}
