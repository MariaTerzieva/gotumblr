gotumblr
========

Description
-----------

A Go Tumblr API v2 Client.

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

For examples on how to work with this API check out the `example_test.go` file in this repository.

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