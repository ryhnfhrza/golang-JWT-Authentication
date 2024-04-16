package service

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ryhnfhrza/golang-JWT-Authentication/model/web"
)

type LoginFormService interface {
	Registration(ctx context.Context, request web.RegistrationRequest) web.LoginFormResponse
	Login(ctx context.Context, request web.LoginRequest) (web.LoginFormResponse,*jwt.Token)
	Logout(ctx context.Context)
}