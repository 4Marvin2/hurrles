package models

type UserLoginBody struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type User struct {
	Id       int64  `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	FullName string `json:"fullName,omitempty"`
	Number   string `json:"number,omitempty"`
}
