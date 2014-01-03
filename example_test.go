package gotumblr_test

import (
		"fmt"
		"github.com/MariaTerzieva/gotumblr"
)

func ExampleInfo(){
	client := gotumblr.NewtumblrRestClient("consumer_key", "consumer_secret", "token", "token_secret", "http://api.tumblr.com")
	info := client.Info()
	fmt.Println(info["response"].(map[string]interface{})["blog"].(map[string]interface{})["name"])
	//Output:
	//mgterzieva
}

func ExampleLikes(){
	client := gotumblr.NewtumblrRestClient("consumer_key", "consumer_secret", "token", "token_secret", "http://api.tumblr.com")
	likes := client.Likes(map[string]int{})
	fmt.Println(likes["response"].(map[string]interface{})["liked_count"])
	//Output:
	//1
}

func ExampleFollowing(){
	client := gotumblr.NewtumblrRestClient("consumer_key", "consumer_secret", "token", "token_secret", "http://api.tumblr.com")
	following := client.Following(map[string]int{})
	fmt.Println(following["response"].(map[string]interface{})["total_blogs"])
	//Output:
	//5
}

func ExampleDashboard(){
	client := gotumblr.NewtumblrRestClient("consumer_key", "consumer_secret", "token", "token_secret", "http://api.tumblr.com")
	dashboard := client.Dashboard(map[string]string{"limit": "1"})
	fmt.Println(dashboard["response"].(map[string]interface{})["blog"].(map[string]interface{})["state"])
	//Output:
	//published	
}

func ExampleTagged(){
	client := gotumblr.NewtumblrRestClient("consumer_key", "consumer_secret", "token", "token_secret", "http://api.tumblr.com")
	tagged := client.Tagged("golang", map[string]string{"limit": "1"})
	fmt.Println(tagged["response"].(map[string]interface{})["blog"].(map[string]interface{})["state"])
	//Output:
	//published
}

func ExamplePosts(){
	client := gotumblr.NewtumblrRestClient("consumer_key", "consumer_secret", "token", "token_secret", "http://api.tumblr.com")
	posts := client.Posts("mgterzieva.tumblr.com", "text", map[string]string{"limit": "1"})
	fmt.Println(posts["response"].(map[string]interface{})["blog"].(map[string]interface{})["total_posts"])
	//Output:
	//2
}

func ExampleBlogInfo(){
	client := gotumblr.NewtumblrRestClient("consumer_key", "consumer_secret", "token", "token_secret", "http://api.tumblr.com")
	blogInfo := client.BlogInfo("mgterzieva.tumblr.com")
	fmt.Println(blogInfo["response"].(map[string]interface{})["blog"].(map[string]interface{})["title"])
	//Output:
	//Testing the GO Tumblr API
}

func ExampleFollowers(){
	client := gotumblr.NewtumblrRestClient("consumer_key", "consumer_secret", "token", "token_secret", "http://api.tumblr.com")
	followers := client.Followers("mgterzieva.tumblr.com", map[string]string{})
	fmt.Println(followers["response"].(map[string]interface{})["total_users"])
	//Output:
	//0
}

func ExampleBlogLikes(){
	client := gotumblr.NewtumblrRestClient("consumer_key", "consumer_secret", "token", "token_secret", "http://api.tumblr.com")
	likes := client.BlogLikes("mgterzieva.tumblr.com", map[string]string{})
	fmt.Println(likes["response"].(map[string]interface{})["liked_count"])
	//Output:
	//1
}

func ExampleQueue(){
	client := gotumblr.NewtumblrRestClient("consumer_key", "consumer_secret", "token", "token_secret", "http://api.tumblr.com")
	queue := client.Queue("mgterzieva.tumblr.com", map[string]string{})
	fmt.Println(queue["response"].(map[string]interface{})["posts"])
	//Output:
	//[]
}

func ExampleDrafts(){
	client := gotumblr.NewtumblrRestClient("consumer_key", "consumer_secret", "token", "token_secret", "http://api.tumblr.com")
	drafts := client.Drafts("mgterzieva.tumblr.com", map[string]string{})
	fmt.Println(drafts["response"].(map[string]interface{})["posts"])
	//Output:
	//[]
}

func ExampleSubmission(){
	client := gotumblr.NewtumblrRestClient("consumer_key", "consumer_secret", "token", "token_secret", "http://api.tumblr.com")
	submission := client.Submission("mgterzieva.tumblr.com", map[string]string{})
	fmt.Println(submission["response"].(map[string]interface{})["posts"])
	//Output:
	//[]
}