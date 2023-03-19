package model

type User struct {
	Username string `json:"username" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Active   bool   `json:"active"`
}
