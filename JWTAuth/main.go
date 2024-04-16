package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/ryhnfhrza/golang-JWT-Authentication/app"
	controllerLoginForm "github.com/ryhnfhrza/golang-JWT-Authentication/controller/LoginFormController"
	controllerProduct "github.com/ryhnfhrza/golang-JWT-Authentication/controller/ProductController"
	"github.com/ryhnfhrza/golang-JWT-Authentication/exception"
	"github.com/ryhnfhrza/golang-JWT-Authentication/helper"
	"github.com/ryhnfhrza/golang-JWT-Authentication/middleware"
	repositoryLoginForm "github.com/ryhnfhrza/golang-JWT-Authentication/repository/LoginFormRepository"
	repositoryProduct "github.com/ryhnfhrza/golang-JWT-Authentication/repository/ProductRepository"
	serviceLoginForm "github.com/ryhnfhrza/golang-JWT-Authentication/service/LoginFormService"
	serviceProduct "github.com/ryhnfhrza/golang-JWT-Authentication/service/ProductService"
)

func main() {
	validate := validator.New()
	db := app.NewDb()
	
	//loginForm
	loginFormRepository := repositoryLoginForm.NewLoginFormRepository()
	loginFormService := serviceLoginForm.NewLoginFormService(loginFormRepository , db ,validate )
	loginFormController := controllerLoginForm.NewLoginFormController(loginFormService)

	//product
	productRepository := repositoryProduct.NewProductRepository()
	productService := serviceProduct.NewProductService(productRepository,db)
	productController := controllerProduct.NewProductController(productService)
	
	router := httprouter.New()

	//loginForm
	router.POST("/registration",loginFormController.Registration)
	router.POST("/login",loginFormController.Login)
	router.GET("/logout", loginFormController.Logout)

	//product
	router.GET("/api/products",middleware.JWTMiddleware(productController.GetAllProduct) )

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr: "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}