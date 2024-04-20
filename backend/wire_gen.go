// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shangsuru/passkey-demo/auth"
	"github.com/shangsuru/passkey-demo/users"
)

// Injectors from wire.go:

func NewServer() (*Server, error) {
	engine := gin.Default()
	userRepository := users.NewUserRepository()
	webAuthn, err := auth.NewWebAuthnAPI()
	if err != nil {
		return nil, err
	}
	webAuthnController := auth.WebAuthnController{
		UserStore:   userRepository,
		WebAuthnAPI: webAuthn,
	}
	server := &Server{
		router:             engine,
		webAuthnController: webAuthnController,
	}
	return server, nil
}
