package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
	"github.com/ryhnfhrza/golang-JWT-Authentication/exception"
	"github.com/ryhnfhrza/golang-JWT-Authentication/util"
)

func JWTMiddleware(next httprouter.Handle)httprouter.Handle{
	return func(writer http.ResponseWriter, request *http.Request,params httprouter.Params) {
		cookie ,err := request.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie{
				panic(exception.NewUnauthorizedError("Token Not Found"))
				
			}
		}

		tokenString := cookie.Value

		claims := &util.JWTClaim{}

		token, err := jwt.ParseWithClaims(tokenString , claims , func(t *jwt.Token) (interface{}, error) {
			return util.JWT_KEY,nil
		})

		if err != nil{
			switch err {
			case jwt.ErrSignatureInvalid:
				exception.WriteUnauthorizedError(writer,err.Error())
				
			case jwt.ErrTokenExpired:
				exception.WriteUnauthorizedError(writer,err.Error())
			default:
				exception.WriteUnauthorizedError(writer,err.Error())
			}
			return
		}

		if !token.Valid{
			panic(exception.NewUnauthorizedError("UNAUTHORIZED"))
		}
		next(writer,request,params)
	}
}