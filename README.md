gotumblr
========

Description
-----------

A Go [Tumblr API](http://www.tumblr.com/docs/en/api/v2) v2 Client.

[![GoDoc](https://godoc.org/github.com/MariaTerzieva/gotumblr?status.png)](https://godoc.org/github.com/MariaTerzieva/gotumblr)

Install gotumblr
----------------

In terminal write `go get github.com/MariaTerzieva/gotumblr`

Running the tests
-----------------

Run the tests with `go test` to check if everything is ok.

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

Then create NewTumblrRestClient with your credentials(consumer key, consumer secret, token, token secret and callback url):

		client := gotumblr.NewTumblrRestClient("consumer_key", "consumer_secret", "token", "token_secret", "callback_url", "http://api.tumblr.com")

Then use the client you just created to get the information you need. Here are some examples with what I got for my account:

		info := client.Info()
		fmt.Println(info.User.Name)
		//Output:
		//mgterzieva

		likes := client.Likes(map[string]string{})
		fmt.Println(likes.Liked_count)
		//Output:
		//63

		following := client.Following(map[string]string{})
		fmt.Println(following.Total_blogs)
		//Output:
		//1

		dashboard := client.Dashboard(map[string]string{"limit": "1"})
		if len(dashboard.Posts) != 0 {
			var base_dashboard_post gotumblr.BasePost
			for i, _ := range dashboard.Posts {
				json.Unmarshal(dashboard.Posts[i], &base_dashboard_post)
				fmt.Println(base_dashboard_post.State)
				//Output:
				//published
			}
		}

		tagged := client.Tagged("golang", map[string]string{"limit": "1"})
		if len(tagged) != 0 {
			var base_tagged_post gotumblr.BasePost
			for i, _ := range tagged {
				json.Unmarshal(tagged[i], &base_tagged_post)
				fmt.Println(base_tagged_post.State)
				//Output:
				//published
			}
		}

		blogname := "mgterzieva.tumblr.com" //this is my blogname. Change this according to your usecase and credentials.
		blogInfo := client.BlogInfo(blogname)
		fmt.Println(blogInfo.Blog.Title)
		//Output:
		//Maria's blog

		followers := client.Followers(blogname, map[string]string{})
		fmt.Println(followers.Total_users)
		//Output:
		//0

		blog_likes := client.BlogLikes(blogname, map[string]string{})
		fmt.Println(blog_likes.Liked_count)
		//Output:
		//63

		queue := client.Queue(blogname, map[string]string{})
		fmt.Println(len(queue.Posts))
		//Output:
		//0

		drafts := client.Drafts(blogname, map[string]string{})
		fmt.Println(len(drafts.Posts))
		//Output:
		//6

		submission := client.Submission(blogname, map[string]string{})
		fmt.Println(len(submission.Posts))
		//Output:
		//0

		avatar := client.Avatar(blogname, 64)
		fmt.Println(avatar.Avatar_url)
		//Output:
		//http://25.media.tumblr.com/avatar_49f49d0b9209_64.png

		other_blogname := "http://thehungergamesmovie.tumblr.com"
		follow := client.Follow(other_blogname)
		fmt.Println(follow)
		//Output:
		//<nil>

		unfollow := client.Unfollow(other_blogname)
		fmt.Println(unfollow)
		//Output:
		//<nil>

		id := "72078164824" //this is the id of a post of mine. Change this according to your usecase.
		//There is an Id field in all of the post object types in this library.
		reblogKey := "6l3e2pGL" //this is the reblogKey of a post of mine. Change this according to your usecase.
		//There is a Reblog_key field in all of the post object types in this library.
		like := client.Like(id, reblogKey)
		fmt.Println(like)
		//Output:
		//<nil>

		unlike := client.Unlike(id, reblogKey)
		fmt.Println(unlike)
		//Output:
		//<nil>

		reblog := client.Reblog(blogname, map[string]string{"id": id, "reblog_key": reblogKey})
		fmt.Println(reblog)
		//Output:
		//<nil>

		state := "draft"
		textPost := client.CreateText(blogname, map[string]string{"body": "Hello happy world!", "state": state})
		fmt.Println(textPost)
		//Output:
		//<nil>

		quote := "A happy heart makes the face cheerful."
		source := "Proverbs 15:13"
		quotePost := client.CreateQuote(blogname, map[string]string{"quote": quote, "source": source, "state": state})
		fmt.Println(quotePost)
		//Output:
		//<nil>

		title := "Follow me on tumblr, guys! :)"
		url := "http://mgterzieva.tumblr.com"
		linkPost := client.CreateLink(blogname, map[string]string{"url": url, "title": title, "state": state})
		fmt.Println(linkPost)
		//Output:
		//<nil>

		conversation := "John Doe: Hi there!\nJane Doe: Hi, John!\nJane Doe: ♥♥♥"
		//separate the tags with commas and don't leave whitespaces around the commas
		tags := "Saint Valentine's day,14th of February,lots of love,xoxo"
		chatPost := client.CreateChatPost(blogname, map[string]string{"conversation": conversation, "tags": tags, "state": state})
		fmt.Println(chatPost)
		//Output:
		//<nil>

		text := "Hello happy world!" //if you are editing a text post
		editPost := client.EditPost(blogname, map[string]string{"id": id, "body": text})
		fmt.Println(editPost)
		//Output:
		//<nil>

		deletePost := client.DeletePost(blogname, id)
		fmt.Println(deletePost)
		//Output:
		//<nil>

		code := `<iframe width="560" height="315" src="//www.youtube.com/embed/uJNvZRAmeqY" frameborder="0" allowfullscreen></iframe>`
		caption := "<b>Mother knows best</b>"
		embedVideo := client.CreateVideo(blogname, map[string]string{"embed": code, "state": state, "caption": caption})
		fmt.Println(embedVideo)
		//Output:
		//<nil>

		song := "https://soundcloud.com/tiffany-alvord-song/the-one-that-got-away-cover-by"
		songPostByURL := client.CreateAudio(blogname, map[string]string{"external_url": song, "state": state})
		fmt.Println(songPostByURL)
		//Output:
		//<nil>

		picture := "http://thumbs.dreamstime.com/z/cute-panda-17976617.jpg"
		photoPostByURL := client.CreatePhoto(blogname, map[string]string{"source": picture, "state": state})
		fmt.Println(photoPostByURL)
		//Output:
		//<nil>

Further information
-------------------

If you still don't get how to work with this package don't worry! :)
Read the [project documentation](https://godoc.org/github.com/MariaTerzieva/gotumblr) or
the [Tumblr API](http://www.tumblr.com/docs/en/api/v2) or
write an e-mail at mgterzieva@abv.bg if these don't help. :)

Contribution
------------
In case you find any issues with this code, use the project's Issues page to report them or send pull requests.

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
