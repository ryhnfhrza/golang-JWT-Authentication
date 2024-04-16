package repository

import (
	"context"
	"database/sql"

	"github.com/ryhnfhrza/golang-JWT-Authentication/model/domain"
)

type LoginFormRepository interface{
	Registration(ctx context.Context, tx *sql.Tx, loginForm domain.LoginForm) domain.LoginForm
	Login(ctx context.Context, tx *sql.Tx, username string) (domain.LoginForm,error)
}