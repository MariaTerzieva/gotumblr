package gotumblr

import "github.com/kurrik/oauth1a"

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
func (tr TumblrRequest) Get(url string, params map[string]string) map[string]interface{}{} {

}

//Makes a POST request to the API, allows for multipart data uploads
//url: the url you are making the request to
//params: all the parameters needed for the request
//files: list of files
func (tr TumblrRequest) Post(url string, params map[string]string, files []string) map[string]interface{}{} {

}

//Parse JSON response.
//content: the content returned from the web request to be pares as JSON
func (tr TumblrRequest) JsonParse(content) map[string]interface{}{} {

}

//Generates and makes a multipart request for data files
//url: the url you are making the request to
//params: all parameters needed for the request
//files: a list of files
func (tr TumblrRequest) PostMultipart(url string, params map[string]string, files []string) map[string]interface{}{} {

}

//Properly encodes the multipart body of the request
//fields: the parameters used in the request
//files: a list of lists containing information about the files
func (tr TumblrRequest) EncodeMultipartFormdata(fields map[string]string, files []string) []string {

}
