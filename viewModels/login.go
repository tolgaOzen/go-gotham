package viewModels

import "gotham/models"

type Login struct {
	AccessToken    string      `json:"access_token"`
	AccessTokenExp int64       `json:"access_token_exp"`
	User           models.User `json:"user"`
}
