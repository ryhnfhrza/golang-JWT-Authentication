package repository

import (
	"context"
	"database/sql"

	"github.com/ryhnfhrza/golang-JWT-Authentication/helper"
	"github.com/ryhnfhrza/golang-JWT-Authentication/model/domain"
)

type ProductRepositoryImpl struct{

}

func NewProductRepository()ProductRepository{
	return &ProductRepositoryImpl{}
}

func(Repository *ProductRepositoryImpl)GetAllProduct( ctx context.Context, tx *sql.Tx) []domain.Product{
	SQL := "select id,name from product"
	rows , err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	
	var products []domain.Product

	for rows.Next(){
		product := domain.Product{}
		err := rows.Scan(&product.Id,&product.Name)
		helper.PanicIfError(err)
		products = append(products,product)
	}
	return products
}