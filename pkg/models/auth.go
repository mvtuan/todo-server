package models

type LoginForm struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
