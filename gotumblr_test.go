package gotumblr

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"reflect"
	"errors"
	"fmt"
)

var (
	//mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	//client is the Tublr client being tested.
	client *TumblrRestClient

	//server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

//setup sets up a test HTTP server along with a gotumblr.TumblrRestClient that is
// configured to talk to that server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() {
	//test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	//tumblr client configured to use test server
	host, _ := url.Parse(server.URL)
	client = NewTumblrRestClient("", "", "", "", "", host.String())
}

//teardown closes the test HTTP server.
func teardown() {
	server.Close()
}

func TestNewTumblrRestClient(t *testing.T) {
	c := NewTumblrRestClient("", "", "", "", "", "http://api.tumblr.com")
	if c.request.host != "http://api.tumblr.com" {
		t.Errorf("New Client host = %v, want http://api.tumblr.com")
	}
}

func TestInfo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/user/info", func(w http.ResponseWriter, r *http.Request) {
			if m := "GET"; m != r.Method {
				t.Errorf("Request method = %v, want %v", r.Method, m)
			}
			fmt.Fprint(w, `{"response": {"user": {"name": "mgterzieva"}}}`)
		})

	info := client.Info().User
	want := UserInfo{Name: "mgterzieva"}
	if !reflect.DeepEqual(info, want) {
		t.Errorf("Info returned %+v, want %+v", info, want)
	}
}

func TestLikes(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/user/likes", func(w http.ResponseWriter, r *http.Request) {
			if m := "GET"; m != r.Method {
				t.Errorf("Request method = %v, want %v", r.Method, m)
			}
			fmt.Fprint(w, `{"response": {"liked_count": 63}}`)
		})

	likes := client.Likes(map[string]string{}).Liked_count
	want := int64(63)
	if likes != want {
		t.Errorf("Likes returned %+v, want %v", likes, want)
	}
}

func TestFollowing(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/user/following", func(w http.ResponseWriter, r *http.Request) {
			if m := "GET"; m != r.Method {
				t.Errorf("Request method = %v, want %v", r.Method, m)
			}
			fmt.Fprint(w, `{"response": {"total_blogs": 1}}`)
		})

	following := client.Following(map[string]string{}).Total_blogs
	want := int64(1)
	if following != want {
		t.Errorf("Following returned %+v, want %v", following, want)
	}
}

func TestDashboard(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/user/dashboard", func(w http.ResponseWriter, r *http.Request) {
			if m := "GET"; m != r.Method {
				t.Errorf("Request method = %v, want %v", r.Method, m)
			}
			fmt.Fprint(w, `{"response": {"posts": [{"type": "photo"}]}}`)
		})

	posts := client.Dashboard(map[string]string{}).Posts
	var post BasePost
	json.Unmarshal(posts[0], &post)
	want := "photo"
	if post.PostType != want {
		t.Errorf("Posts returned %+v, want %v", post.PostType, want)
	}
}

func TestTagged(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/tagged", func(w http.ResponseWriter, r *http.Request) {
			if m := "GET"; m != r.Method {
				t.Errorf("Request method = %v, want %v", r.Method, m)
			}
			fmt.Fprint(w, `{"response": [{"format": "html"}]}`)
		})

	posts := client.Tagged("golang", map[string]string{})
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

	mux.HandleFunc("/v2/blog/mgterzieva/posts/html", func(w http.ResponseWriter, r *http.Request) {
			if m := "GET"; m != r.Method {
				t.Errorf("Request method = %v, want %v", r.Method, m)
			}
			fmt.Fprint(w, `{"response": {"blog": {"description": "none"}, "total_posts": 8}}`)
		})

	data := client.Posts("mgterzieva", "html", map[string]string{})
	want := "none"
	if data.Blog.Description != want {
		t.Errorf("Description returned %+v, want %v", data.Blog.Description, want)
	}
	expected := int64(8)
	if data.Total_posts != expected {
		t.Errorf("Total_posts returned %+v, expected %v", data.Total_posts, expected)
	}
}

func TestAvatar(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/blog/mgterzieva/avatar/64", func(w http.ResponseWriter, r *http.Request) {
			if m := "GET"; m != r.Method {
				t.Errorf("Request method = %v, want %v", r.Method, m)
			}
			fmt.Fprint(w, `{"response": {"avatar_url": "http://cool-pic.jpg"}}`)
		})

	avatar := client.Avatar("mgterzieva", 64).Avatar_url
	want := "http://cool-pic.jpg"
	if avatar != want {
		t.Errorf("Avatar returned %+v, want %+v", avatar, want)
	}
}

func TestBlogInfo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/blog/mgterzieva/info", func(w http.ResponseWriter, r *http.Request) {
			if m := "GET"; m != r.Method {
				t.Errorf("Request method = %v, want %v", r.Method, m)
			}
			fmt.Fprint(w, `{"response": {"blog": {"updated": 1392218146, "ask": false, "ask_anon": false}}}`)
		})

	info := client.BlogInfo("mgterzieva").Blog
	want := BlogInfo{Updated: 1392218146, Ask: false, Ask_anon: false}
	if info != want {
		t.Errorf("BlogInfo returned %+v, want %+v", info, want)
	}
}

func TestFollowers(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/blog/mgterzieva/followers", func(w http.ResponseWriter, r *http.Request) {
			if m := "GET"; m != r.Method {
				t.Errorf("Request method = %v, want %v", r.Method, m)
			}
			fmt.Fprint(w, `{"response": {"total_users": 0, "users": []}}`)
		})

	followers := client.Followers("mgterzieva", map[string]string{})
	want := FollowersResponse{Total_users: 0, Users: []User{}}
	if !reflect.DeepEqual(followers, want) {
		t.Errorf("Followers returned %+v, want %+v", followers, want)
	}
}

func TestBlogLikes(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/blog/mgterzieva/likes", func(w http.ResponseWriter, r *http.Request) {
			if m := "GET"; m != r.Method {
				t.Errorf("Request method = %v, want %v", r.Method, m)
			}
			fmt.Fprint(w, `{"response": {"liked_posts": [], "liked_count": 0}}`)
		})

	likes := client.BlogLikes("mgterzieva", map[string]string{})
	want := LikesResponse{Liked_posts: []json.RawMessage{}, Liked_count: 0}
	if !reflect.DeepEqual(likes, want) {
		t.Errorf("BlogLikes returned %+v, want %+v", likes, want)
	}
}

func TestQueue(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/blog/mgterzieva/posts/queue", func(w http.ResponseWriter, r *http.Request) {
			if m := "GET"; m != r.Method {
				t.Errorf("Request method = %v, want %v", r.Method, m)
			}
			fmt.Fprint(w, `{"response": {"posts": []}}`)
		})

	queue := client.Queue("mgterzieva", map[string]string{})
	want := DraftsResponse{Posts: []json.RawMessage{}}
	if !reflect.DeepEqual(queue, want) {
		t.Errorf("Queue returned %+v, want %+v", queue, want)
	}
}

func TestDrafts(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/blog/mgterzieva/posts/draft", func(w http.ResponseWriter, r *http.Request) {
			if m := "GET"; m != r.Method {
				t.Errorf("Request method = %v, want %v", r.Method, m)
			}
			fmt.Fprint(w, `{"response": {"posts": []}}`)
		})

	drafts := client.Drafts("mgterzieva", map[string]string{})
	want := DraftsResponse{Posts: []json.RawMessage{}}
	if !reflect.DeepEqual(drafts, want) {
		t.Errorf("Drafts returned %+v, want %+v", drafts, want)
	}
}

func TestSubmission(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/blog/mgterzieva/posts/submission", func(w http.ResponseWriter, r *http.Request) {
			if m := "GET"; m != r.Method {
				t.Errorf("Request method = %v, want %v", r.Method, m)
			}
			fmt.Fprint(w, `{"response": {"posts": []}}`)
		})

	submission := client.Submission("mgterzieva", map[string]string{})
	want := DraftsResponse{Posts: []json.RawMessage{}}
	if !reflect.DeepEqual(submission, want) {
		t.Errorf("Submission returned %+v, want %+v", submission, want)
	}
}

func TestFollow(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/user/follow", func(w http.ResponseWriter, r *http.Request) {
			if m := "POST"; m != r.Method {
				t.Errorf("Request method = %v, want %v", r.Method, m)
			}
			fmt.Fprint(w, `{"meta": {"status":404, "msg": "Not Found"}}`)
		})

	follow := client.Follow("thehungergames")
	want := errors.New("Not Found")
	if !reflect.DeepEqual(follow, want) {
		t.Errorf("Follow returned %+v, want %+v", follow, want)
	}
}

func TestUnfollow(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/user/unfollow", func(w http.ResponseWriter, r *http.Request) {
			if m := "POST"; m != r.Method {
				t.Errorf("Request method = %v, want %v", r.Method, m)
			}
			fmt.Fprint(w, `{"meta": {"status":404, "msg": "Not Found"}}`)
		})

	unfollow := client.Unfollow("thehungergames")
	want := errors.New("Not Found")
	if !reflect.DeepEqual(unfollow, want) {
		t.Errorf("Unfollow returned %+v, want %+v", unfollow, want)
	}
}

func TestLike(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/user/like", func(w http.ResponseWriter, r *http.Request) {
			if m := "POST"; m != r.Method {
				t.Errorf("Request method = %v, want %v", r.Method, m)
			}
			fmt.Fprint(w, `{"meta": {"status":200, "msg": "OK"}}`)
		})

	like := client.Like("75195127536", "kLXwhQ19")
	if like != nil {
		t.Errorf("Like returned %+v, want %+v", like, nil)
	}
}

func TestUnlike(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/user/unlike", func(w http.ResponseWriter, r *http.Request) {
			if m := "POST"; m != r.Method {
				t.Errorf("Request method = %v, want %v", r.Method, m)
			}
			fmt.Fprint(w, `{"meta": {"status":200, "msg": "OK"}}`)
		})

	unlike := client.Unlike("75195127536", "kLXwhQ19")
	if unlike != nil {
		t.Errorf("Unlike returned %+v, want %+v", unlike, nil)
	}
}

func TestCreatePhoto(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/blog/mgterzieva/post", func(w http.ResponseWriter, r *http.Request) {
			if m := "POST"; m != r.Method {
				t.Errorf("Request method = %v, want %v", r.Method, m)
			}
			fmt.Fprint(w, `{"meta": {"status":400, "msg": "Bad Request"}}`)
		})

	post_photo := client.CreatePhoto("mgterzieva", map[string]string{"state": "draft"})
	want := errors.New("Bad Request")
	if !reflect.DeepEqual(post_photo, want) {
		t.Errorf("CreatePhoto returned %+v, want %+v", post_photo, want)
	}
}

func TestCreateText(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/blog/mgterzieva/post", func(w http.ResponseWriter, r *http.Request) {
			if m := "POST"; m != r.Method {
				t.Errorf("Request method = %v, want %v", r.Method, m)
			}
			fmt.Fprint(w, `{"meta": {"status": 201, "msg": "Created"}}`)
		})

	post_text := client.CreateText("mgterzieva", map[string]string{"body": "Hello, hello!"})
	if post_text != nil {
		t.Errorf("CreateText returned %+v, want %+v", post_text, nil)
	}
}

func TestCreateQuote(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/blog/mgterzieva/post", func(w http.ResponseWriter, r *http.Request) {
			if m := "POST"; m != r.Method {
				t.Errorf("Request method = %v, want %v", r.Method, m)
			}
			fmt.Fprint(w, `{"meta": {"status": 201, "msg": "Created"}}`)
		})
	quote := "You can complain because roses have thorns, or you can rejoice because thorns have roses."
	source := "Ziggy"
	post_quote := client.CreateQuote("mgterzieva", map[string]string{"source": source, "quote": quote})
	if post_quote != nil {
		t.Errorf("CreateQuote returned %+v, want %+v", post_quote, nil)
	}
}

func TestCreateChatPost(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/blog/mgterzieva/post", func(w http.ResponseWriter, r *http.Request) {
			if m := "POST"; m != r.Method {
				t.Errorf("Request method = %v, want %v", r.Method, m)
			}
			fmt.Fprint(w, `{"meta": {"status": 400, "msg": "Bad Request"}}`)
		})

	post_discussion := client.CreateChatPost("mgterzieva", map[string]string{})
	want := errors.New("Bad Request")
	if !reflect.DeepEqual(post_discussion, want) {
		t.Errorf("CreateChatPost returned %+v, want %+v", post_discussion, nil)
	}
}

func TestCreateAudio(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/blog/mgterzieva/post", func(w http.ResponseWriter, r *http.Request) {
			if m := "POST"; m != r.Method {
				t.Errorf("Request method = %v, want %v", r.Method, m)
			}
			fmt.Fprint(w, `{"meta": {"status": 201, "msg": "Created"}}`)
		})

	post_song := client.CreateAudio("mgterzieva", map[string]string{"external_url": "http://coolsongs.com/song"})
	if post_song != nil {
		t.Errorf("CreateAudio returned %+v, want %+v", post_song, nil)
	}
}

func TestCreateVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/blog/mgterzieva/post", func(w http.ResponseWriter, r *http.Request) {
			if m := "POST"; m != r.Method {
				t.Errorf("Request method = %v, want %v", r.Method, m)
			}
			fmt.Fprint(w, `{"meta": {"status": 201, "msg": "Created"}}`)
		})
	code := `<iframe width="560" height="315" src="//www.videos.com/embed/uMNGkgsgaB" frameborder="0" allowfullscreen></iframe>`
	post_video := client.CreateVideo("mgterzieva", map[string]string{"embed": code})
	if post_video != nil {
		t.Errorf("CreateVideo returned %+v, want %+v", post_video, nil)
	}
}

func TestReblog(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/blog/mgterzieva/post/reblog", func(w http.ResponseWriter, r *http.Request) {
			if m := "POST"; m != r.Method {
				t.Errorf("Request method = %v, want %v", r.Method, m)
			}
			fmt.Fprint(w, `{"meta": {"status": 400, "msg": "Bad Request"}}`)
		})

	reblog := client.Reblog("mgterzieva", map[string]string{"id": "7161981", "reblog_key": "blah"})
	want := errors.New("Bad Request")
	if !reflect.DeepEqual(reblog, want) {
		t.Errorf("Reblog returned %+v, want %+v", reblog, nil)
	}
}

func TestDeletePost(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/blog/mgterzieva/post/delete", func(w http.ResponseWriter, r *http.Request) {
			if m := "POST"; m != r.Method {
				t.Errorf("Request method = %v, want %v", r.Method, m)
			}
			fmt.Fprint(w, `{"meta": {"status": 400, "msg": "Bad Request"}}`)
		})

	delete := client.DeletePost("mgterzieva", "")
	want := errors.New("Bad Request")
	if !reflect.DeepEqual(delete, want) {
		t.Errorf("DeletePost returned %+v, want %+v", delete, want)
	}
}

func TestEditPost(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/blog/mgterzieva/post/edit", func(w http.ResponseWriter, r *http.Request) {
			if m := "POST"; m != r.Method {
				t.Errorf("Request method = %v, want %v", r.Method, m)
			}
			fmt.Fprint(w, `{"meta": {"status": 400, "msg": "Bad Request"}}`)
		})

	edit := client.EditPost("mgterzieva", map[string]string{})
	want := errors.New("Bad Request")
	if !reflect.DeepEqual(edit, want) {
		t.Errorf("EditPost returned %+v, want %+v", edit, want)
	}
}