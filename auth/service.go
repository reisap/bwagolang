package auth

import (
	"github.com/dgrijalva/jwt-go"
)


type Service interface {
	GenerateToken(userID int) (string,error)

}

type jwtService struct {

}

func NewJwtService() *jwtService{
	return &jwtService{}
}

var SECRET_KEY = []byte("anakLUCUbgt#!@")

func (s *jwtService)GenerateToken(userID int) (string,error){
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	verifyToken,err := token.SignedString(SECRET_KEY)

	if err != nil {
		return verifyToken,err
	}

	return verifyToken,nil

}