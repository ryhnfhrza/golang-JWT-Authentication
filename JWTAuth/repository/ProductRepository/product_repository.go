package repository

import (
	"context"
	"database/sql"

	"github.com/ryhnfhrza/golang-JWT-Authentication/model/domain"
)

type ProductRepository interface {
	GetAllProduct( ctx context.Context, tx *sql.Tx) []domain.Product
}