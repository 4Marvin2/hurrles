package models

//easyjson:json
type UserCredentials struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

//easyjson:json
type User struct {
	Id       uint64 `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	FullName string `json:"fullName,omitempty"`
	Number   string `json:"number,omitempty"`
}

//easyjson:json
type Session struct {
	Cookie string
	UserID uint64
}
