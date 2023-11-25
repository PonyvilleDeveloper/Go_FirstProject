package entity

import "time"

type User struct {
	UserId    uint32
	FIO       string
	Email     string
	DateBirth time.Time
	Password  string
	Login     string
}
