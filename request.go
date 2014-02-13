//A Go Tumblr API v2 Client.

package gotumblr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/kurrik/oauth1a"
)

//Make queries to the Tumblr API through TumblrRequest.
type TumblrRequest struct {
	service    *oauth1a.Service
	userConfig *oauth1a.UserConfig
	host       string
	apiKey     string
}

//Initializes the TumblrRequest.
//consumerKey is the consumer key of your Tumblr Application.
//consumerSecret is the consumer secret of your Tumblr Application.
//callbackUrl is the callback URL of your Tumblr Application.
//oauthToken is the user specific token, received from the /access_token endpoint.
//oauthSecret is the user specific secret, received from the /access_token endpoint.
//host is the host that you are tryng to send information to (e.g. http://api.tumblr.com).
func NewTumblrRequest(consumerKey, consumerSecret, oauthToken, oauthSecret, callbackUrl, host string) *TumblrRequest {
	service := &oauth1a.Service{
		RequestURL:   "http://www.tumblr.com/oauth/request_token",
		AuthorizeURL: "http://www.tumblr.com/oauth/authorize",
		AccessURL:    "http://www.tumblr.com/oauth/access_token",
		ClientConfig: &oauth1a.ClientConfig{
			ConsumerKey:    consumerKey,
			ConsumerSecret: consumerSecret,
			CallbackURL:    callbackUrl,
		},
		Signer: new(oauth1a.HmacSha1Signer),
	}
	userConfig := oauth1a.NewAuthorizedConfig(oauthToken, oauthSecret)
	return &TumblrRequest{service, userConfig, host, consumerKey}
}

//Make a GET request to the API with properly formatted parameters.
//requestUrl: the url you are making the request to.
//params: the parameters needed for the request.
func (tr *TumblrRequest) Get(requestUrl string, params map[string]string) CompleteResponse {
	fullUrl := tr.host + requestUrl
	if len(params) != 0 {
		values := url.Values{}
		for key, value := range params {
			values.Set(key, value)
		}
		fullUrl = fullUrl + "?" + values.Encode()
	}
	httpRequest, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		fmt.Println(err)
	}
	tr.service.Sign(httpRequest, tr.userConfig)
	var httpResponse *http.Response
	httpClient := new(http.Client)
	httpResponse, err2 := httpClient.Do(httpRequest)
	if err2 != nil {
		fmt.Println(err2)
	}
	defer httpResponse.Body.Close()
	body, err3 := ioutil.ReadAll(httpResponse.Body)
	if err3 != nil {
		fmt.Println(err3)
	}
	return tr.JSONParse(body)
}

//Makes a POST request to the API, allows for multipart data uploads.
//requestUrl: the url you are making the request to.
//params: all the parameters needed for the request.
func (tr *TumblrRequest) Post(requestUrl string, params map[string]string) CompleteResponse {
	full_url := tr.host + requestUrl
	values := url.Values{}
	for key, value := range params {
		values.Set(key, value)
	}
	httpRequest, err := http.NewRequest("POST", full_url, strings.NewReader(values.Encode()))
	if err != nil {
		fmt.Println(err)
	}
	httpRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	tr.service.Sign(httpRequest, tr.userConfig)
	var httpResponse *http.Response
	httpClient := new(http.Client)
	httpResponse, err2 := httpClient.Do(httpRequest)
	if err2 != nil {
		fmt.Println(err2)
	}
	defer httpResponse.Body.Close()
	body, err3 := ioutil.ReadAll(httpResponse.Body)
	if err3 != nil {
		fmt.Println(err3)
	}
	return tr.JSONParse(body)
}

//Parse JSON response.
//content: the content returned from the web request to be parsed as JSON.
func (tr *TumblrRequest) JSONParse(content []byte) CompleteResponse {
	var data CompleteResponse
	err := json.Unmarshal(content, &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}
