package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ryhnfhrza/golang-JWT-Authentication/exception"
	"github.com/ryhnfhrza/golang-JWT-Authentication/helper"
	"github.com/ryhnfhrza/golang-JWT-Authentication/model/domain"
	"github.com/ryhnfhrza/golang-JWT-Authentication/model/web"
	repository "github.com/ryhnfhrza/golang-JWT-Authentication/repository/LoginFormRepository"
	"github.com/ryhnfhrza/golang-JWT-Authentication/util"
	utils "github.com/ryhnfhrza/golang-JWT-Authentication/util"
	"golang.org/x/crypto/bcrypt"
)

type LoginFormServiceImpl struct {
	LoginFormRepository repository.LoginFormRepository
	Db                  *sql.DB
	validate            *validator.Validate
}

func NewLoginFormService(loginFormRepository repository.LoginFormRepository, db *sql.DB, Validate *validator.Validate) LoginFormService {
	return &LoginFormServiceImpl{
		LoginFormRepository: loginFormRepository,
		Db: db,
		validate:        Validate,
	}
}

func(service *LoginFormServiceImpl)Registration(ctx context.Context, request web.RegistrationRequest) web.LoginFormResponse{
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	tx,err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	hashPassword,err := utils.HashPassword(request.Password)
	helper.PanicIfError(err)

	loginForm := domain.LoginForm{
		Username: request.Username,
		Email: request.Email,
		Password: hashPassword,
	}

	loginForm = service.LoginFormRepository.Registration(ctx,tx,loginForm)

	return helper.ToLoginFormResponse(loginForm)
}

func(service *LoginFormServiceImpl)Login(ctx context.Context, request web.LoginRequest) (web.LoginFormResponse,*jwt.Token){
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	tx,err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	
	
	login,err := service.LoginFormRepository.Login(ctx,tx,request.Username)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}
	
	
	err = bcrypt.CompareHashAndPassword([]byte(login.Password),[]byte(request.Password))
	if err != nil{
		panic(exception.NewUnauthorizedError("Password is incorrect"))
	}
	
	expTime := time.Now().Add(time.Hour * 1)
	claims := &util.JWTClaim{
		Username: request.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "github.com/ryhnfhrza",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	
	
	return helper.ToLoginFormResponse(login),tokenAlgo
}

func(service *LoginFormServiceImpl)Logout(ctx context.Context){

}
