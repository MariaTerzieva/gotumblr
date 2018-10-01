package gotumblr

import (
	"encoding/json"
	"net/url"
	"testing"
)

func TestGet(t *testing.T) {
	setup()
	defer teardown()

	res := `{"meta": {"status": 200}, "response": {"user": {"name": "mgterzieva"}}}`
	handleFunc("/v2/user/info", "GET", res, url.Values{}, t)
	data, err := client.Get("/v2/user/info", url.Values{})
	if err != nil {
		t.Fatal(err)
	}

	expectedMeta := MetaInfo{Status: 200}
	var response UserInfoResponse
	if err := json.Unmarshal(data.Response, &response); err != nil {
		t.Fatal(err)
	}
	expectedName := "mgterzieva"

	if data.Meta != expectedMeta {
		t.Errorf("Get returned %+v, want %+v", data.Meta, expectedMeta)
	}
	if response.User.Name != expectedName {
		t.Errorf("Get returned %v, want %v", response.User.Name, expectedName)
	}
}

func TestPost(t *testing.T) {
	setup()
	defer teardown()

	response := `{"meta": {"status":404, "msg": "Not Found"}}`

	handleFunc("/v2/user/follow", "POST", response, url.Values{
		"url": []string{"thehungergames"},
	}, t)
	data, _ := client.Post("/v2/user/follow", url.Values{
		"url": []string{"thehungergames"},
	})
	expectedMeta := MetaInfo{Msg: "Not Found", Status: 404}
	if data.Meta != expectedMeta {
		t.Errorf("Post returned %+v, want %+v", data.Meta, expectedMeta)
	}
}
