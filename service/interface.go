package service

import (
	"go-napi/auth"
	"go-napi/connection"
)

type CommandHandler struct {
	ai auth.AuthInterface
	di connection.DatabaseInterface
}

func NewHandler(ai auth.AuthInterface, di connection.DatabaseInterface) *CommandHandler {
	return &CommandHandler{ai, di}
}
