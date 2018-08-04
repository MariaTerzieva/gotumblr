package gotumblr

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the Tublr client being tested.
	client *Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

// setup sets up a test HTTP server along with a gotumblr.TumblrRestClient that is
// configured to talk to that server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// tumblr client configured to use test server
	host, _ := url.Parse(server.URL)
	client = New("", "", "", "", SetHost(host.String()))
}

// teardown closes the test HTTP server.
func teardown() {
	server.Close()
}

// checks the request parameters
func checkParameters(request *http.Request, parameters url.Values, t *testing.T) {
	request.ParseForm()
	for key := range parameters {
		if request.Form.Get(key) != parameters.Get(key) {
			t.Errorf("%v should be %v", key, parameters.Get(key))
		}
	}
}

func handleFunc(url, method, response string, parameters url.Values, t *testing.T) {
	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		checkParameters(r, parameters, t)
		if m := method; m != r.Method {
			t.Errorf("Request method = %v, want %v", r.Method, m)
		}
		fmt.Fprint(w, response)
	})
}

func TestNew(t *testing.T) {
	c := New("", "", "", "", SetHost("http://api.tumblr.com"))
	if c.host != "http://api.tumblr.com" {
		t.Errorf("New Client host = %v, want http://api.tumblr.com", c.host)
	}
	if c.apiKey != "" {
		t.Errorf("New Client host = %v, want the empty string", c.apiKey)
	}
}

func TestInfo(t *testing.T) {
	setup()
	defer teardown()

	body := `{"meta": {"status": 200}, "response": {"user": {"name": "mgterzieva"}}}`
	handleFunc("/v2/user/info", "GET", body, url.Values{}, t)

	info, err := client.Info()
	if err != nil {
		t.Fatal(err)
	}
	want := UserInfo{Name: "mgterzieva"}
	if !reflect.DeepEqual(info.User, want) {
		t.Errorf("Info returned %+v, want %+v", info.User, want)
	}
}

func TestLikes(t *testing.T) {
	setup()
	defer teardown()

	body := `{"meta": {"status": 200}, "response": {"liked_count": 63}}`
	handleFunc("/v2/user/likes", "GET", body, url.Values{}, t)

	likes, err := client.Likes(url.Values{})
	if err != nil {
		t.Fatal(err)
	}
	want := int64(63)
	if likes.LikedCount != want {
		t.Errorf("Likes returned %+v, want %v", likes.LikedCount, want)
	}
}

func TestFollowing(t *testing.T) {
	setup()
	defer teardown()

	handleFunc("/v2/user/following", "GET", `{"meta": {"status": 200}, "response": {"total_blogs": 1}}`, url.Values{}, t)

	following, err := client.Following(url.Values{})
	if err != nil {
		t.Fatal(err)
	}
	want := int64(1)
	if following.TotalBlogs != want {
		t.Errorf("Following returned %+v, want %v", following.TotalBlogs, want)
	}
}

func TestDashboard(t *testing.T) {
	setup()
	defer teardown()

	handleFunc("/v2/user/dashboard", "GET", `{"meta": {"status": 200}, "response": {"posts": [{"type": "photo"}]}}`, url.Values{}, t)

	posts, err := client.Dashboard(url.Values{})
	if err != nil {
		t.Fatal(err)
	}
	var post BasePost
	json.Unmarshal(posts.Posts[0], &post)
	want := "photo"
	if post.PostType != want {
		t.Errorf("Posts returned %+v, want %v", post.PostType, want)
	}
}

func TestTagged(t *testing.T) {
	setup()
	defer teardown()

	res := `{"meta": {"status": 200}, "response": [{"format": "html"}]}`
	handleFunc("/v2/tagged", "GET", res, url.Values{}, t)

	posts, err := client.Tagged("golang", url.Values{})
	if err != nil {
		t.Fatal(err)
	}
	var post BasePost
	json.Unmarshal(posts[0], &post)
	want := "html"
	if post.Format != want {
		t.Errorf("Posts returned %+v, want %v", post.Format, want)
	}
}

func TestPosts(t *testing.T) {
	setup()
	defer teardown()

	res := `{"meta": {"status": 200}, "response": {"blog": {"description": "none"}, "total_posts": 8}}`
	handleFunc("/v2/blog/mgterzieva/posts/html", "GET", res, url.Values{}, t)

	data, err := client.Posts("mgterzieva", "html", url.Values{})
	if err != nil {
		t.Fatal(err)
	}
	want := "none"
	if data.Blog.Description != want {
		t.Errorf("Description returned %+v, want %v", data.Blog.Description, want)
	}
	expected := int64(8)
	if data.TotalPosts != expected {
		t.Errorf("TotalPosts returned %+v, expected %v", data.TotalPosts, expected)
	}
}

func TestAvatar(t *testing.T) {
	setup()
	defer teardown()

	res := `{"meta": {"status": 200}, "response": {"avatar_url": "http://cool-pic.jpg"}}`
	handleFunc("/v2/blog/mgterzieva/avatar/64", "GET", res, url.Values{}, t)

	avatar, err := client.Avatar("mgterzieva", 64)
	if err != nil {
		t.Fatal(err)
	}
	want := "http://cool-pic.jpg"
	if avatar.AvatarURL != want {
		t.Errorf("Avatar returned %+v, want %+v", avatar.AvatarURL, want)
	}
}

func TestBlogInfo(t *testing.T) {
	setup()
	defer teardown()

	res := `{"meta": {"status": 200}, "response": {"blog": {"updated": 1392218146, "ask": false, "ask_anon": false}}}`
	handleFunc("/v2/blog/mgterzieva/info", "GET", res, url.Values{}, t)

	info, err := client.BlogInfo("mgterzieva")
	if err != nil {
		t.Fatal(err)
	}
	want := BlogInfo{Updated: 1392218146, Ask: false, AskAnon: false}
	if info.Blog != want {
		t.Errorf("BlogInfo returned %+v, want %+v", info.Blog, want)
	}
}

func TestFollowers(t *testing.T) {
	setup()
	defer teardown()

	res := `{"meta": {"status": 200}, "response": {"total_users": 0, "users": []}}`
	handleFunc("/v2/blog/mgterzieva/followers", "GET", res, url.Values{}, t)

	followers, err := client.Followers("mgterzieva", url.Values{})
	if err != nil {
		t.Fatal(err)
	}
	want := FollowersResponse{TotalUsers: 0, Users: []User{}}
	if !reflect.DeepEqual(followers, want) {
		t.Errorf("Followers returned %+v, want %+v", followers, want)
	}
}

func TestBlogLikes(t *testing.T) {
	setup()
	defer teardown()

	res := `{"meta": {"status": 200}, "response": {"liked_posts": [], "liked_count": 0}}`
	handleFunc("/v2/blog/mgterzieva/likes", "GET", res, url.Values{}, t)

	likes, err := client.BlogLikes("mgterzieva", url.Values{})
	if err != nil {
		t.Fatal(err)
	}
	want := LikesResponse{LikedPosts: []json.RawMessage{}, LikedCount: 0}
	if !reflect.DeepEqual(likes, want) {
		t.Errorf("BlogLikes returned %+v, want %+v", likes, want)
	}
}

func TestQueue(t *testing.T) {
	setup()
	defer teardown()

	res := `{"meta": {"status": 200}, "response": {"posts": []}}`
	handleFunc("/v2/blog/mgterzieva/posts/queue", "GET", res, url.Values{}, t)

	queue, err := client.Queue("mgterzieva", url.Values{})
	if err != nil {
		t.Fatal(err)
	}
	want := DraftsResponse{Posts: []json.RawMessage{}}
	if !reflect.DeepEqual(queue, want) {
		t.Errorf("Queue returned %+v, want %+v", queue, want)
	}
}

func TestDrafts(t *testing.T) {
	setup()
	defer teardown()

	res := `{"meta": {"status": 200}, "response": {"posts": []}}`
	handleFunc("/v2/blog/mgterzieva/posts/draft", "GET", res, url.Values{}, t)

	drafts, err := client.Drafts("mgterzieva", url.Values{})
	if err != nil {
		t.Fatal(err)
	}
	want := DraftsResponse{Posts: []json.RawMessage{}}
	if !reflect.DeepEqual(drafts, want) {
		t.Errorf("Drafts returned %+v, want %+v", drafts, want)
	}
}

func TestSubmission(t *testing.T) {
	setup()
	defer teardown()

	res := `{"meta": {"status": 200}, "response": {"posts": []}}`
	handleFunc("/v2/blog/mgterzieva/posts/submission", "GET", res, url.Values{}, t)

	submission, err := client.Submission("mgterzieva", url.Values{})
	if err != nil {
		t.Fatal(err)
	}
	want := DraftsResponse{Posts: []json.RawMessage{}}
	if !reflect.DeepEqual(submission, want) {
		t.Errorf("Submission returned %+v, want %+v", submission, want)
	}
}

func TestFollow(t *testing.T) {
	setup()
	defer teardown()

	res := `{"meta": {"status": 404, "msg": "Not Found"}}`
	handleFunc("/v2/user/follow", "POST", res, url.Values{
		"url": []string{"thehungergames"},
	}, t)

	follow := client.Follow("thehungergames")
	want := errors.New("Not Found")
	if !reflect.DeepEqual(follow, want) {
		t.Errorf("Follow returned %+v, want %+v", follow, want)
	}
}

func TestUnfollow(t *testing.T) {
	setup()
	defer teardown()

	res := `{"meta": {"status":404, "msg": "Not Found"}}`
	handleFunc("/v2/user/unfollow", "POST", res, url.Values{
		"url": []string{"thehungergames"},
	}, t)

	unfollow := client.Unfollow("thehungergames")
	want := errors.New("Not Found")
	if !reflect.DeepEqual(unfollow, want) {
		t.Errorf("Unfollow returned %+v, want %+v", unfollow, want)
	}
}

func TestLike(t *testing.T) {
	setup()
	defer teardown()

	res := `{"meta": {"status":200, "msg": "OK"}}`
	handleFunc("/v2/user/like", "POST", res, url.Values{
		"id":         []string{"75195127536"},
		"reblog_key": []string{"kLXwhQ19"},
	}, t)

	like := client.Like("75195127536", "kLXwhQ19")
	if like != nil {
		t.Errorf("Like returned %+v, want %+v", like, nil)
	}
}

func TestUnlike(t *testing.T) {
	setup()
	defer teardown()

	res := `{"meta": {"status": 200, "msg": "OK"}}`
	handleFunc("/v2/user/unlike", "POST", res, url.Values{
		"id":         []string{"75195127536"},
		"reblog_key": []string{"kLXwhQ19"},
	}, t)

	unlike := client.Unlike("75195127536", "kLXwhQ19")
	if unlike != nil {
		t.Errorf("Unlike returned %+v, want %+v", unlike, nil)
	}
}

func TestCreatePhoto(t *testing.T) {
	setup()
	defer teardown()

	res := `{"meta": {"status": 400, "msg": "Bad Request"}}`
	handleFunc("/v2/blog/mgterzieva/post", "POST", res, url.Values{
		"state": []string{"draft"},
	}, t)

	err := client.CreatePhoto("mgterzieva", url.Values{
		"state": []string{"draft"},
	})
	want := errors.New("Bad Request")
	if !reflect.DeepEqual(err, want) {
		t.Errorf("CreatePhoto returned %+v, want %+v", err, want)
	}
}

func TestCreateText(t *testing.T) {
	setup()
	defer teardown()

	res := `{"meta": {"status": 201, "msg": "Created"}}`
	handleFunc("/v2/blog/mgterzieva/post", "POST", res, url.Values{
		"body": []string{"Hello, hello!"},
	}, t)

	err := client.CreateText("mgterzieva", url.Values{
		"body": []string{"Hello, hello!"},
	})
	if err != nil {
		t.Errorf("CreateText returned %+v, want %+v", err, nil)
	}
}

func TestCreateQuote(t *testing.T) {
	setup()
	defer teardown()

	quote := "You can complain because roses have thorns, or you can rejoice because thorns have roses."
	source := "Ziggy"
	res := `{"meta": {"status": 201, "msg": "Created"}}`
	handleFunc("/v2/blog/mgterzieva/post", "POST", res, url.Values{
		"source": []string{source},
		"quote":  []string{quote},
	}, t)

	err := client.CreateQuote("mgterzieva", url.Values{
		"source": []string{source},
		"quote":  []string{quote},
	})
	if err != nil {
		t.Errorf("CreateQuote returned %+v, want %+v", err, nil)
	}
}

func TestCreateChatPost(t *testing.T) {
	setup()
	defer teardown()

	res := `{"meta": {"status": 400, "msg": "Bad Request"}}`
	handleFunc("/v2/blog/mgterzieva/post", "POST", res, url.Values{}, t)

	err := client.CreateChatPost("mgterzieva", url.Values{})
	want := errors.New("Bad Request")
	if !reflect.DeepEqual(err, want) {
		t.Errorf("CreateChatPost returned %+v, want %+v", err, nil)
	}
}

func TestCreateAudio(t *testing.T) {
	setup()
	defer teardown()

	res := `{"meta": {"status": 201, "msg": "Created"}}`
	handleFunc("/v2/blog/mgterzieva/post", "POST", res, url.Values{
		"externalURL": []string{"http://coolsongs.com/song"},
	}, t)

	err := client.CreateAudio("mgterzieva", url.Values{
		"externalURL": []string{"http://coolsongs.com/song"},
	})
	if err != nil {
		t.Errorf("CreateAudio returned %+v, want %+v", err, nil)
	}
}

func TestCreateVideo(t *testing.T) {
	setup()
	defer teardown()
	code := `<iframe width="560" height="315" src="//www.videos.com/embed/uMNGkgsgaB" frameborder="0" allowfullscreen></iframe>`
	res := `{"meta": {"status": 201, "msg": "Created"}}`
	handleFunc("/v2/blog/mgterzieva/post", "POST", res, url.Values{
		"embed": []string{code},
	}, t)

	err := client.CreateVideo("mgterzieva", url.Values{
		"embed": []string{code},
	})
	if err != nil {
		t.Errorf("CreateVideo returned %+v, want %+v", err, nil)
	}
}

func TestReblog(t *testing.T) {
	setup()
	defer teardown()

	res := `{"meta": {"status": 400, "msg": "Bad Request"}}`
	handleFunc("/v2/blog/mgterzieva/post/reblog", "POST", res, url.Values{
		"id":         []string{"7161981"},
		"reblog_key": []string{"blah"},
	}, t)

	reblog := client.Reblog("mgterzieva", url.Values{
		"id":         []string{"7161981"},
		"reblog_key": []string{"blah"},
	})
	want := errors.New("Bad Request")
	if !reflect.DeepEqual(reblog, want) {
		t.Errorf("Reblog returned %+v, want %+v", reblog, nil)
	}
}

func TestDeletePost(t *testing.T) {
	setup()
	defer teardown()

	res := `{"meta": {"status": 400, "msg": "Bad Request"}}`
	handleFunc("/v2/blog/mgterzieva/post/delete", "POST", res, url.Values{
		"id": []string{""},
	}, t)

	delete := client.DeletePost("mgterzieva", "")
	want := errors.New("Bad Request")
	if !reflect.DeepEqual(delete, want) {
		t.Errorf("DeletePost returned %+v, want %+v", delete, want)
	}
}

func TestEditPost(t *testing.T) {
	setup()
	defer teardown()

	res := `{"meta": {"status": 400, "msg": "Bad Request"}}`
	handleFunc("/v2/blog/mgterzieva/post/edit", "POST", res, url.Values{}, t)

	edit := client.EditPost("mgterzieva", url.Values{})
	want := errors.New("Bad Request")
	if !reflect.DeepEqual(edit, want) {
		t.Errorf("EditPost returned %+v, want %+v", edit, want)
	}
}
