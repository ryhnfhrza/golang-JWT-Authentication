package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ryhnfhrza/golang-JWT-Authentication/helper"
	"github.com/ryhnfhrza/golang-JWT-Authentication/model/domain"
)

type LoginFormRepositoryImpl struct{

}

func NewLoginFormRepository()LoginFormRepository{
	return &LoginFormRepositoryImpl{}
}

func(Repository *LoginFormRepositoryImpl)Registration(ctx context.Context, tx *sql.Tx, LoginForm domain.LoginForm) domain.LoginForm{
	SQL := "insert into login_form (username,password,email) values (?,?,?)"
	_,err := tx.ExecContext(ctx,SQL,LoginForm.Username,LoginForm.Password,LoginForm.Email)
	helper.PanicIfError(err)

	return LoginForm
}

func(Repository *LoginFormRepositoryImpl)Login(ctx context.Context, tx *sql.Tx, username string) (domain.LoginForm,error){
	SQL := "select username,password,email from login_form where username = ?"
	rows , err := tx.QueryContext(ctx, SQL, username)
	helper.PanicIfError(err)
	
	loginForm := domain.LoginForm{}
	if rows.Next(){
		err := rows.Scan(&loginForm.Username,&loginForm.Password,&loginForm.Email)
		helper.PanicIfError(err)
		return loginForm,nil
	}else{
		return loginForm,errors.New("username " + username + " not found")
	}
}





