package models

import (
	"time"
)

type User struct {
	Acct      string `gorm:"primaryKey`
	Pwd       string
	FullName  string `gorm:"column:fullname"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
