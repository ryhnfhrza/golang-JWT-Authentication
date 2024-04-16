package service

import (
	"context"

	"github.com/ryhnfhrza/golang-JWT-Authentication/model/web"
)

type LoginFormService interface {
	Registration(ctx context.Context, request web.RegistrationRequest) web.LoginFormResponse
	Login(ctx context.Context, request web.LoginRequest) web.LoginFormResponse
	Logout(ctx context.Context)
}