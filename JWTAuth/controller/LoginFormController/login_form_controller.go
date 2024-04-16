package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type LoginFormController interface {
	Registration(writer http.ResponseWriter, request *http.Request,params httprouter.Params)
	Login(writer http.ResponseWriter, request *http.Request,params httprouter.Params)
	Logout(writer http.ResponseWriter, request *http.Request,params httprouter.Params)
}