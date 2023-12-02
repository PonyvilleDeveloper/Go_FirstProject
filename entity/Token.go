package entity

import "time"

type Token struct {
	TokenId uint32    `json:"-" gorm:"primaryKey"`
	UserId  uint32    `json:"UserId"`
	Token   string    `json:"Token"`
	Expires time.Time `json:"Expired"`
}

func NewToken(user_id uint32) Token {
	var res Token
	res.TokenId = 0
	res.UserId = user_id
	res.Token = "dwsvdsvdhgvsadgvs"
	res.Expires = time.Now()
	return res
}
