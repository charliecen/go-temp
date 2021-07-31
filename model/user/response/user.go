package response

import "go-temp/model/user"

type UserResponse struct {
	User user.User `json:"user"`
}

type LoginResponse struct {
	User      user.User `json:"user"`
	Token     string    `json:"token"`
	ExpiresAt int64     `json:"expiresAt"`
}
