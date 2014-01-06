gotumblr
========

Description
-----------

A Go Tumblr API v2 Client.

Dependencies
------------

This package uses [kurrik's oauth1a package](https://github.com/kurrik/oauth1a).
In order to use the gotumblr package, you have to install the oauth1a package like this:

`go get github.com/kurrik/oauth1a`

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

For examples how to work with the API see the `example_test.go` file in this repository.
