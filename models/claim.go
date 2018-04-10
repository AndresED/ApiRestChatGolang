package models

//JWT =>go get github.com/dgrijalva/jwt-go
import jwt "github.com/dgrijalva/jwt-go"

// Claim estructura que define los datos de nuestro modelo Claim
type Claim struct {
	User `json:"user"`
	jwt.StandardClaims
}
