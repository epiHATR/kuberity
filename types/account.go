package types

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Exception struct {
	Message string `json:"message"`
}
