# gotumblr

[![GoDoc](https://godoc.org/github.com/BakeRolls/gotumblr?status.svg)](https://godoc.org/github.com/BakeRolls/gotumblr)
[![Go Report Card](https://goreportcard.com/badge/github.com/BakeRolls/gotumblr)](https://goreportcard.com/report/github.com/BakeRolls/gotumblr)

A Go [Tumblr API](http://www.tumblr.com/docs/en/api/v2) v2 Client. This is a fork of [gotumblr](https://github.com/MariaTerzieva/gotumblr) by [Maria Terzieva](https://github.com/MariaTerzieva).

```go
func main() {
	t := gotumblr.New("consumerKey", "consumerSecret", "accessToken", "accessTokenSecret")
	info, err := t.Info()
	if err != nil {
		return log.Fatal(err)
	}
	info, err := t.Info()
	if err != nil {
		return log.Fatal(err)
	}
	fmt.Pritnf("Hello %s!", info.User)
}
```

# Caching

You can use `gotumblr.SetClient` and `gotumblr.SetHeaders` to cache responses with something like [httpcache](https://github.com/gregjones/httpcache).

```go
func main() {
	t = gotumblr.New("consumerKey", "consumerSecret", "accessToken", "accessTokenSecret",
		gotumblr.SetClient(httpcache.NewTransport(diskcache.New("cache"))),
		gotumblr.SetHeaders(map[string]string{"cache-control": "max-stale=60"}),
	)

	// ...
}
```
