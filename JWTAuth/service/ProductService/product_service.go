package service

import (
	"context"

	"github.com/ryhnfhrza/golang-JWT-Authentication/model/web"
)

type ProductService interface {
	GetAllProduct(ctx context.Context)[]web.ProductResponse
}