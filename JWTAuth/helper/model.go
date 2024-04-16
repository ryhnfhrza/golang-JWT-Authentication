package helper

import (
	"github.com/ryhnfhrza/golang-JWT-Authentication/model/domain"
	"github.com/ryhnfhrza/golang-JWT-Authentication/model/web"
)

func ToLoginFormResponse(loginForm domain.LoginForm)web.LoginFormResponse{
	return web.LoginFormResponse{
		Username: loginForm.Username,
		Email: loginForm.Email,
	}

}

func ToLoginFormResponses(loginForm [] domain.LoginForm)[]web.LoginFormResponse{
	var loginFormResponses []web.LoginFormResponse
	for _,lf := range loginForm{
		loginFormResponses = append(loginFormResponses, ToLoginFormResponse(lf))
	}
	return loginFormResponses
}

func ToProductResponse(product domain.Product)web.ProductResponse{
	return web.ProductResponse{
		Id: product.Id,
		Name: product.Name,
	}

}

func ToProductResponses(product [] domain.Product)[]web.ProductResponse{
	var productResponses []web.ProductResponse
	for _,prdct := range product{
		productResponses = append(productResponses, ToProductResponse(prdct))
	}
	return productResponses
}