package gotumblr

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"reflect"
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