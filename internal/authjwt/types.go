package authjwt

import "github.com/dgrijalva/jwt-go/v4"

type CustomClaims struct {
	Login string
	jwt.StandardClaims
}
