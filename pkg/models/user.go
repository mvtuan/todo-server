package models

import "time"

type User struct {
	//gorm.Model

	CommonModel `json:",inline"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
	FullName    string `json:"fullName,omitempty"`

	Token        string     `json:"token,omitempty" gorm:"-"`
	LastLoggedIn *time.Time `json:"lastLoggedIn,omitempty"`
}
