package gotumblr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Get makes a GET request to the API with properly formatted parameters.
// requestURL: the url you are making the request to.
// values: the parameters needed for the request.
func (c *Client) Get(requestURL string, values url.Values) (CompleteResponse, error) {
	fullURL := c.host + requestURL + "?" + values.Encode()
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return CompleteResponse{}, err
	}
	for k, v := range c.headers {
		req.Header.Set(k, v)
	}
	if err := c.service.Sign(req, c.userConfig); err != nil {
		return CompleteResponse{}, err
	}
	res, err := c.client.Do(req)
	if err != nil {
		return CompleteResponse{}, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return CompleteResponse{}, err
	}
	var data CompleteResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return CompleteResponse{}, err
	}
	if data.Meta.Status < 200 || data.Meta.Status >= 300 {
		return data, fmt.Errorf(data.Meta.Msg)
	}
	return data, nil
}

// Post makes a POST request to the API, allows for multipart data uploads.
// requestURL: the url you are making the request to.
// values: all the parameters needed for the request.
func (c *Client) Post(requestURL string, values url.Values) (CompleteResponse, error) {
	fullURL := c.host + requestURL
	req, err := http.NewRequest("POST", fullURL, strings.NewReader(values.Encode()))
	if err != nil {
		return CompleteResponse{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for k, v := range c.headers {
		req.Header.Set(k, v)
	}
	c.service.Sign(req, c.userConfig)
	res, err := c.client.Do(req)
	if err != nil {
		return CompleteResponse{}, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return CompleteResponse{}, err
	}
	var data CompleteResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return CompleteResponse{}, err
	}
	if data.Meta.Status < 200 || data.Meta.Status >= 300 {
		return data, fmt.Errorf(data.Meta.Msg)
	}
	return data, nil
}
