package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ProductController interface {
	GetAllProduct(writer http.ResponseWriter, request *http.Request,params httprouter.Params)
}