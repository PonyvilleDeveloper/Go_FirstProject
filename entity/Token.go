package entity

import "time"

type Token struct {
	TokenId uint32    `json:"-" gorm:"primaryKey"`
	UserId  uint32    `json:"UserId"`
	Token   string    `json:"Token"`
	Expired time.Time `json:"Expired"`
}
