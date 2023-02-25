package models

import (
	"time"
)

type User struct {
	Acct        string    `json:"acct" gorm:"primaryKey`
	Pwd         string    `json:"pwd"`
	FullName    string    `json:"fullname"`
	DateCreated time.Time `json:"created_at"`
	DateUpdated time.Time `json:"updated_at"`
}
