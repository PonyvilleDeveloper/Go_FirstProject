package entity

import "time"

type User struct {
	UserId    uint64
	FIO       string
	Email     string
	DateBirth time.Time
}