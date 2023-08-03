package user

import "time"

type User struct {
	ID           int
	Name         string
	Nim          string
	Email        string
	Password     string
	Avatar       string
	RoleID       bool
	Division     string
	NoHP         string
	AlasanDaftar string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
