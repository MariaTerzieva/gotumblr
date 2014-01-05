package gotumblr

import (
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/kurrik/oauth1a"
)

//Make queries to the Tumblr API through TumblrRequest
type TumblrRequest struct {
	service *oauth1a.Service
	userConfig *oauth1a.UserConfig
	host string
}

//Initializes the TumblrRequest.
//consumerKey is the consumer key of your Tumblr Application
//consumerSecret is the consumer secret of your Tumblr Application
//callbackUrl is the callback URL of your Tumblr Application
//oauthToken is the user specific token, received from the /access_token endpoint
//oauthSecret is the user specific secret, received from the /access_token endpoint
//host is the host that you are tryng to send information to (e.g. http://api.tumblr.com)
func NewTumblrRequest(consumerKey, consumerSecret, oauthToken, oauthSecret, callbackUrl, host string) *TumblrRequest {
	service := &oauth1a.Service{
		RequestURL:   "http://www.tumblr.com/oauth/request_token",
		AuthorizeURL: "http://www.rumblr.com/oauth/authorize",
		AccessURL:    "http://www.tumblr.com/oauth/access_token",
		ClientConfig: &oauth1a.ClientConfig{
			ConsumerKey:    consumerKey,
			ConsumerSecret: consumerSecret,
			CallbackURL:    callbackUrl,
	    },
		Signer: new(oauth1a.HmacSha1Signer),
	}
	userConfig := oauth1a.NewAuthorizedConfig(oauthToken, oauthSecret)
	return &TumblrRequest{service, userConfig, host}
}

//Make a GET request to the API with properly formatted parameters
//url: the url you are making the request to
//params: the parameters needed for the request 
func (tr *TumblrRequest) Get(url string, params map[string]string) map[string]interface{} {
	full_url := tr.host + url
	if len(params) != 0 {
		values := url.Values{}
		for key, value := range params {
			values.Set(key, value)
			full_url = full_url + "?" + values.Encode() 
		}
	}
	httpRequest, err := http.NewRequest("GET", full_url, nil)
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
	return tr.JsonParse(body)
}

//Makes a POST request to the API, allows for multipart data uploads
//url: the url you are making the request to
//params: all the parameters needed for the request
//files: list of files
func (tr *TumblrRequest) Post(url string, params map[string]string, files []string) map[string]interface{} {

}

//Parse JSON response.
//content: the content returned from the web request to be parsed as JSON
func (tr *TumblrRequest) JsonParse(content []byte) map[string]interface{} {
	data := map[string]interface{}{}
	err := json.Unmarshal(content, &data)
	if err != nil {
		fmt.Println(err)
	}
	ok_statuses := []int{200, 201, 301}
	status := data["meta"].(map[string]interface{})["status"]
	for _, ok_status := range ok_statuses {
		if status == ok_status {
			return data["response"].(map[string]interface{})
		}
	}
	return data
}

//Generates and makes a multipart request for data files
//url: the url you are making the request to
//params: all parameters needed for the request
//files: a list of files
func (tr *TumblrRequest) PostMultipart(url string, params map[string]string, files []string) map[string]interface{} {

}

//Properly encodes the multipart body of the request
//fields: the parameters used in the request
//files: a list of lists containing information about the files
func (tr *TumblrRequest) EncodeMultipartFormdata(fields map[string]string, files []string) (string, string) {

}
