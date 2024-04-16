package controller

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
	"github.com/ryhnfhrza/golang-JWT-Authentication/helper"
	"github.com/ryhnfhrza/golang-JWT-Authentication/model/web"
	service "github.com/ryhnfhrza/golang-JWT-Authentication/service/LoginFormService"
	"github.com/ryhnfhrza/golang-JWT-Authentication/util"
)

type LoginFormControllerImpl struct {
	LoginFormService service.LoginFormService
	
}

func NewLoginFormController(loginFormservice service.LoginFormService)LoginFormController{
	return &LoginFormControllerImpl{
		LoginFormService: loginFormservice,
	}
}

func(controller *LoginFormControllerImpl)Registration(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	LoginFormCreateRequest := web.RegistrationRequest{}
	helper.ReadFromRequestBody(request,&LoginFormCreateRequest)

	LoginFormResponse := controller.LoginFormService.Registration(request.Context(),LoginFormCreateRequest)
	webResponse := web.WebResponse{
		Code: http.StatusCreated,
		Status: "CREATED",
		Data: LoginFormResponse,
	}

	helper.WriteToResponseBody(writer,webResponse)
}

func(controller *LoginFormControllerImpl)Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	LoginRequest := web.LoginRequest{}
	helper.ReadFromRequestBody(request,&LoginRequest)
	
	LoginFormResponse := controller.LoginFormService.Login(request.Context(),LoginRequest)
	
	expTime := time.Now().Add(time.Hour * 1)
	claims := &util.JWTClaim{
		Username: LoginRequest.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "github.com/ryhnfhrza",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	token,err := tokenAlgo.SignedString(util.JWT_KEY)
	if err != nil{
		webResponse := web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
	}

	http.SetCookie(writer,&http.Cookie{
		Name: "token",
		Path: "/",
		Value: token,
		HttpOnly: true,
	})

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   LoginFormResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}


func(controller *LoginFormControllerImpl)Logout(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	http.SetCookie(writer ,& http.Cookie{
		Name: "token",
		Path: "/",
		Value: "",
		HttpOnly: true,
		MaxAge: -1,
	})
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "Ok",
		Data: "Logout",
	}

	helper.WriteToResponseBody(writer,webResponse)

}

