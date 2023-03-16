package types

type AuthResponse struct {
	AccessToken string `json:"accesstoken"`
	UserInfo    User   `json:"userinfo"`
}
