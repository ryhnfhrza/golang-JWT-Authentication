package service

import (
	"context"
	"database/sql"

	"github.com/ryhnfhrza/golang-JWT-Authentication/helper"
	"github.com/ryhnfhrza/golang-JWT-Authentication/model/web"
	repository "github.com/ryhnfhrza/golang-JWT-Authentication/repository/ProductRepository"
)

type ProductServiceImpl struct {
	Product repository.ProductRepository
	Db                  *sql.DB
}

func NewProductService(product repository.ProductRepository, db *sql.DB) ProductService {
	return &ProductServiceImpl{
		Product: product,
		Db: db,
	}
}

func(Service *ProductServiceImpl)GetAllProduct(ctx context.Context)[]web.ProductResponse{
	tx,err := Service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Product := Service.Product.GetAllProduct(ctx,tx)

	return helper.ToProductResponses(Product)
}