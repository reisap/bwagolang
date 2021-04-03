package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)


type Service interface {
	GenerateToken(userID int) (string,error)
	ValidateToken(token string) (*jwt.Token,error)

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

func (s *jwtService)ValidateToken(validateToken string) (*jwt.Token,error){
	token,err := jwt.Parse(validateToken,func(token *jwt.Token)(interface{},error){
		 _, ok := token.Method.(*jwt.SigningMethodHMAC)
		 if !ok {
		 	return nil,errors.New("Invalid Token")
		 }

		 return []byte(SECRET_KEY),nil
	})

	if err != nil {
		return token,err
	}

	return token,nil
}