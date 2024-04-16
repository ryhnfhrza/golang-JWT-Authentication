package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ryhnfhrza/golang-JWT-Authentication/helper"
	"github.com/ryhnfhrza/golang-JWT-Authentication/model/web"
	service "github.com/ryhnfhrza/golang-JWT-Authentication/service/ProductService"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
	
}

func NewProductController(productService service.ProductService)ProductController{
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func(controller *ProductControllerImpl)GetAllProduct(writer http.ResponseWriter, request *http.Request,params httprouter.Params){
	productResponse := controller.ProductService.GetAllProduct(request.Context())
	webResponse := web.WebResponse{
		Code: http.StatusCreated,
		Status: "CREATED",
		Data: productResponse,
	}

	helper.WriteToResponseBody(writer,webResponse)
}
