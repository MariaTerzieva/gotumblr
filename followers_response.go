package gotumblr

type FollowersResponse struct {
	TotalUsers int64  `json:"total_users"`
	Users      []User `json:"users"`
}
