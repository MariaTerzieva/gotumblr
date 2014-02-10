gotumblr
========

Description
-----------

A Go [Tumblr API](http://www.tumblr.com/docs/en/api/v2) v2 Client.

Install gotumblr
----------------

In terminal write `go get github.com/MariaTerzieva/gotumblr`

Using the package
-----------------

To use this package in your projects do this (after install):

`import "github.com/MariaTerzieva/gotumblr"`

You are going to need a consumer key, consumer secret, callback URL, token and token secret.
You can get the consumer key, consumer secret and callback URL by registering a Tumblr application.
You can do so by clicking [here](http://www.tumblr.com/oauth/apps).
The token and token secret you can get by using OAUTH.
If you want to use this package just for your own tumblr account, after registering the application,
click on the Explore API option and allow it access to your Tumblr account. You are going to see a tab "Show keys".
Click on it and you will get your token and token secret but if you want, you can also obtain them using OAUTH.

Examples
--------

First import the package in your project as shown above.

Then create NewTublrRestClient with your credentials(consumer key, consumer secret, token, token secret and callback url):

		client := gotumblr.NewTumblrRestClient("consumer_key", "consumer_secret", "token", "token_secret", "callback_url", "http://api.tumblr.com")

Then use the client you just created to get the information you need. Here are some examples:

		info := client.Info()
		fmt.Println(info["response"].(map[string]interface{})["blog"].(map[string]interface{})["name"])
		//Output:
		//the name of the user's blog (e.g. blogname in blogname.tumblr.com)

		likes := client.Likes(map[string]string{})
		fmt.Println(likes["response"].(map[string]interface{})["liked_count"])
		//Output:
		//the count of the posts the user has liked

		following := client.Following(map[string]string{})
		fmt.Println(following["response"].(map[string]interface{})["total_blogs"])
		//Output:
		//the number of the blogs the user is following

		dashboard := client.Dashboard(map[string]string{"limit": "1"})
		fmt.Println(dashboard["response"].(map[string]interface{})["blog"].(map[string]interface{})["state"])
		//Output:
		//published

		tagged := client.Tagged("golang", map[string]string{"limit": "1"})
		fmt.Println(tagged["response"].(map[string]interface{})["blog"].(map[string]interface{})["state"])
		//Output:
		//published

		blogname := "example.tumblr.com" //please change the blogname according to your credentials
		blogInfo := client.BlogInfo(blogname)
		fmt.Println(blogInfo["response"].(map[string]interface{})["blog"].(map[string]interface{})["title"])
		//Output:
		//the title of the blog

		blogname := "example.tumblr.com" //please change the blogname according to your credentials
		followers := client.Followers(blogname, map[string]string{})
		fmt.Println(followers["response"].(map[string]interface{})["total_users"])
		//Output:
		//the number of all followers of the blog

		blogname := "example.tumblr.com" //please change the blogname according to your credentials 
		likes := client.BlogLikes(blogname, map[string]string{})
		fmt.Println(likes["response"].(map[string]interface{})["liked_count"])
		//Output:
		//the number of all blog likes

		blogname := "example.tumblr.com" //please change the blogname according to your credentials 
		queue := client.Queue(blogname, map[string]string{})
		fmt.Println(queue["response"].(map[string]interface{})["posts"])
		//Output:
		//an interface of all posts in the queue


		blogname := "example.tumblr.com" //please change the blogname according to your credentials
		drafts := client.Drafts(blogname, map[string]string{})
		fmt.Println(drafts["response"].(map[string]interface{})["posts"])
		//Output:
		//an interface of all posts in the drafts section

		blogname := "example.tumblr.com" //please change the blogname according to your credentials
		submission := client.Submission(blogname, map[string]string{})
		fmt.Println(submission["response"].(map[string]interface{})["posts"])
		//Output:
		//an interface of all posts in the submissions section

		blogname := "example.tumblr.com" //please change the blogname according to your credentials
		avatar := client.Avatar(blogname, 64)
		fmt.Println(avatar["meta"].(map[string]interface{})["status"])
		//Output:
		//301

		blogname := "mgterzieva.tumblr.com"
		follow := client.Follow(blogname)
		fmt.Println(follow["meta"].(map[string]interface{})["status"])
		//Output:
		//200

		blogname := "mgterzieva.tumblr.com"
		unfollow := client.Unfollow(blogname)
		fmt.Println(unfollow["meta"].(map[string]interface{})["status"])
		//Output:
		//200

		id := "72078164824"
		reblogKey := "6l3e2pGL"
		like := client.Like(id, reblogKey)
		fmt.Println(like["meta"].(map[string]interface{})["status"])
		//Output:
		//200

		id := "72078164824"
		reblogKey := "6l3e2pGL"
		unlike := client.Unlike(id, reblogKey)
		fmt.Println(unlike["meta"].(map[string]interface{})["status"])
		//Output:
		//200

		blogname := "example.tumblr.com" //please change the blogname according to your credentials
		id := "72078164824"
		reblogKey := "6l3e2pGL"
		reblog := client.Reblog(blogname, map[string]string{"id": id, "reblog_key": reblogKey})
		fmt.Println(reblog["meta"].(map[string]interface{})["status"])
		//Output:
		//201

		blogname := "example.tumblr.com" //please change the blogname according to your credentials
		textPost := client.CreateText(blogname, map[string]string{"body": "Hello world!"})
		fmt.Println(textPost["meta"].(map[string]interface{})["status"])
		//Output:
		//201

		blogname := "example.tumblr.com" //please change the blogname according to your credentials
		quotePost := client.CreateQuote(blogname, map[string]string{"quote": "A happy heart makes the face cheerful."})
		fmt.Println(quotePost["meta"].(map[string]interface{})["status"])
		//Output:
		//201

		blogname := "example.tumblr.com" //please change the blogname according to your credentials
		linkPost := client.CreateLink(blogname, map[string]string{"url": "http://mgterzieva.tumblr.com"})
		fmt.Println(linkPost["meta"].(map[string]interface{})["status"])
		//Output:
		//201

		blogname := "example.tumblr.com" //please change the blogname according to your credentials
		conversation := "John Doe: Hi there!\nJane Doe: Hi!"
		chatPost := client.CreateChatPost(blogname, map[string]string{"conversation": conversation})
		fmt.Println(chatPost["meta"].(map[string]interface{})["status"])
		//Output:
		//201

		blogname := "example.tumblr.com" //please change the blogname according to your credentials
		id := "72078164824" //please change the id according to the blogpost
		text := "Hello happy world!" //if you are editing a text post
		editPost := client.EditPost(blogname, map[string]string{"id": id, "text": text})
		fmt.Println(editPost["meta"].(map[string]interface{})["status"])
		//Output:
		//200

		blogname := "example.tumblr.com" //please change the blogname according to your credentials
		id := "72078164824" //please change the id according to the blogpost you want to delete
		deletePost := client.DeletePost(blogname, id)
		fmt.Println(deletePost["meta"].(map[string]interface{})["status"])
		//Output:
		//200

		blogname := "example.tumblr.com" //please change the blogname according to your credentials
		code := `<iframe width="560" height="315" src="//www.youtube.com/embed/T5OEg5paXK0" frameborder="0" allowfullscreen></iframe>`
		embedVideo := client.CreateVideo(blogname, map[string]string{"embed": code})
		fmt.Println(embedVideo["meta"].(map[string]interface{})["status"])
		//Output:
		//201

		blogname := "example.tumblr.com" //please change the blogname according to your credentials
		song := "https://soundcloud.com/tiffany-alvord-song/the-one-that-got-away-cover-by"
		songPostByURL := client.CreateAudio(blogname, map[string]string{"external_url": song})
		fmt.Println(songPostByURL["meta"].(map[string]interface{})["status"])
		//Output:
		//201

		blogname := "example.tumblr.com" //please change the blogname according to your credentials
		picture := "http://herrickshighlander.com/wp-content/uploads/2014/01/aladdin_jasmine_carpet1.jpg"
		photoPostByURL := client.CreatePhoto(blogname, map[string]string{"source": picture})
		fmt.Println(photoPostByURL["meta"].(map[string]interface{})["status"])
		//Output:
		//201

License
-------

Copyright 2014 Maria Terzieva


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.