package gotumblr_test

import (
	"fmt"
	"github.com/MariaTerzieva/gotumblr"
)

func ExampleInfo() {
	client := gotumblr.NewtumblrRestClient("consumer_key", "consumer_secret", "token", "token_secret", "callbackurl", "http://api.tumblr.com")
	info := client.Info()
	fmt.Println(info["response"].(map[string]interface{})["blog"].(map[string]interface{})["name"])
	//Output:
	//the name of the user's blog here(e.g. blogname in blogname.tumblr.com)
}

func ExampleLikes() {
	client := gotumblr.NewtumblrRestClient("consumer_key", "consumer_secret", "token", "token_secret", "callbackurl", "http://api.tumblr.com")
	likes := client.Likes(map[string]string{})
	fmt.Println(likes["response"].(map[string]interface{})["liked_count"])
	//Output:
	//the count of the posts the user has liked here
}

func ExampleFollowing() {
	client := gotumblr.NewtumblrRestClient("consumer_key", "consumer_secret", "token", "token_secret", "callbackurl", "http://api.tumblr.com")
	following := client.Following(map[string]string{})
	fmt.Println(following["response"].(map[string]interface{})["total_blogs"])
	//Output:
	//the number of the blogs the user is following here
}

func ExampleDashboard() {
	client := gotumblr.NewtumblrRestClient("consumer_key", "consumer_secret", "token", "token_secret", "callbackurl", "http://api.tumblr.com")
	dashboard := client.Dashboard(map[string]string{"limit": "1"})
	fmt.Println(dashboard["response"].(map[string]interface{})["blog"].(map[string]interface{})["state"])
	//Output:
	//published
}

func ExampleTagged() {
	client := gotumblr.NewtumblrRestClient("consumer_key", "consumer_secret", "token", "token_secret", "callbackurl", "http://api.tumblr.com")
	tagged := client.Tagged("golang", map[string]string{"limit": "1"})
	fmt.Println(tagged["response"].(map[string]interface{})["blog"].(map[string]interface{})["state"])
	//Output:
	//published
}

func ExamplePosts() {
	blogname := "example.tumblr.com"
	client := gotumblr.NewtumblrRestClient("consumer_key", "consumer_secret", "token", "token_secret", "callbackurl", "http://api.tumblr.com")
	posts := client.Posts(blogname, "text", map[string]string{"limit": "1"})
	fmt.Println(posts["response"].(map[string]interface{})["blog"].(map[string]interface{})["total_posts"])
	//Output:
	//the number of all blog posts here; don't forget to change blogname according to your credentials!
}

func ExampleBlogInfo() {
	blogname := "example.tumblr.com"
	client := gotumblr.NewtumblrRestClient("consumer_key", "consumer_secret", "token", "token_secret", "callbackurl", "http://api.tumblr.com")
	blogInfo := client.BlogInfo(blogname)
	fmt.Println(blogInfo["response"].(map[string]interface{})["blog"].(map[string]interface{})["title"])
	//Output:
	//the title of the blog here; don't forget to change blogname according to your credentials!
}

func ExampleFollowers() {
	blogname := "example.tumblr.com"
	client := gotumblr.NewtumblrRestClient("consumer_key", "consumer_secret", "token", "token_secret", "callbackurl", "http://api.tumblr.com")
	followers := client.Followers(blogname, map[string]string{})
	fmt.Println(followers["response"].(map[string]interface{})["total_users"])
	//Output:
	//the number of all followers of the blog; don't forget to change blogname according to your credentials!
}

func ExampleBlogLikes() {
	blogname := "example.tumblr.com"
	client := gotumblr.NewtumblrRestClient("consumer_key", "consumer_secret", "token", "token_secret", "callbackurl", "http://api.tumblr.com")
	likes := client.BlogLikes(blogname, map[string]string{})
	fmt.Println(likes["response"].(map[string]interface{})["liked_count"])
	//Output:
	//the number of all blog likes here; don't forget to change blogname according to your credentials!
}

func ExampleQueue() {
	blogname := "example.tumblr.com"
	client := gotumblr.NewtumblrRestClient("consumer_key", "consumer_secret", "token", "token_secret", "callbackurl", "http://api.tumblr.com")
	queue := client.Queue(blogname, map[string]string{})
	fmt.Println(queue["response"].(map[string]interface{})["posts"])
	//Output:
	//an interface of all posts in the queue;  don't forget to change blogname according to your credentials!
}

func ExampleDrafts() {
	blogname := "example.tumblr.com"
	client := gotumblr.NewtumblrRestClient("consumer_key", "consumer_secret", "token", "token_secret", "callbackurl", "http://api.tumblr.com")
	drafts := client.Drafts(blogname, map[string]string{})
	fmt.Println(drafts["response"].(map[string]interface{})["posts"])
	//Output:
	//an interface of all posts in the drafts section;  don't forget to change blogname according to your credentials!
}

func ExampleSubmission() {
	blogname := "example.tumblr.com"
	client := gotumblr.NewtumblrRestClient("consumer_key", "consumer_secret", "token", "token_secret", "callbackurl", "http://api.tumblr.com")
	submission := client.Submission(blogname, map[string]string{})
	fmt.Println(submission["response"].(map[string]interface{})["posts"])
	//Output:
	//an interface of all posts in the submissions section;  don't forget to change blogname according to your credentials!
}

func ExampleAvatar() {
	blogname := "example.tumblr.com"
	client := gotumblr.NewtumblrRestClient("consumer_key", "consumer_secret", "token", "token_secret", "callbackurl", "http://api.tumblr.com")
	avatar := client.Avatar(blogname, 64)
	fmt.Println(avatar["meta"].(map[string]interface{})["status"])
	//Output:
	//301
}

func ExampleFollow() {
	blogname := "mgterzieva.tumblr.com"
	client := gotumblr.NewtumblrRestClient("consumer_key", "consumer_secret", "token", "token_secret", "callbackurl", "http://api.tumblr.com")
	follow := client.Follow(blogname)
	fmt.Println(follow["meta"].(map[string]interface{})["status"])
	//Output:
	//200
}

func ExampleUnfollow() {
	blogname := "mgterzieva.tumblr.com"
	client := gotumblr.NewtumblrRestClient("consumer_key", "consumer_secret", "token", "token_secret", "callbackurl", "http://api.tumblr.com")
	unfollow := client.Unfollow(blogname)
	fmt.Println(unfollow["meta"].(map[string]interface{})["status"])
	//Output:
	//200
}

func ExampleLike() {
	id := 72078164824
	reblogKey := "6l3e2pGL"
	client := gotumblr.NewtumblrRestClient("consumer_key", "consumer_secret", "token", "token_secret", "callbackurl", "http://api.tumblr.com")
	like := client.Like(id, reblogKey)
	fmt.Println(like["meta"].(map[string]interface{})["status"])
	//Output:
	//200
}

func ExampleUnlike() {
	id := 72078164824
	reblogKey := "6l3e2pGL"
	client := gotumblr.NewtumblrRestClient("consumer_key", "consumer_secret", "token", "token_secret", "callbackurl", "http://api.tumblr.com")
	unlike := client.Unlike(id, reblogKey)
	fmt.Println(unlike["meta"].(map[string]interface{})["status"])
	//Output:
	//200
}

func ExampleReblog() {
	blogname := "example.tumblr.com"
	id := "72078164824"
	reblogKey := "6l3e2pGL"
	client := gotumblr.NewtumblrRestClient("consumer_key", "consumer_secret", "token", "token_secret", "callbackurl", "http://api.tumblr.com")
	reblog := client.Reblog(blogname, map[string]string{"id": id, "reblog_key": reblogKey})
	fmt.Println(reblog["meta"].(map[string]interface{})["status"])
	//Output:
	//201
}
