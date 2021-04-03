package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User,error)
	Login (input LoginInput) (User,error)
	IsCheckEmail (input CheckEmailInput) (bool,error)
	SaveAvatar (ID int,fileLocation string) (User,error)
	GetUserByID(ID int) (User,error)
}

type userService struct {
	repository Repository
}

func NewService(repository Repository) *userService{
	return &userService{repository}
}

func (s *userService) RegisterUser(input RegisterUserInput) (User,error){
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	PasswordHash,err := bcrypt.GenerateFromPassword([]byte(input.Password),bcrypt.MinCost)
	if err != nil {
		return user,err
	}
	user.PasswordHash = string(PasswordHash)
	user.Role = "user"

	newUser,err := s.repository.Save(user)
	if err != nil {
		return newUser,err
	}

	return newUser,nil

}

func (s *userService)Login (input LoginInput) (User,error){
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user,err
	}
	if user.ID == 0 {
		return user,errors.New("No user found on this email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash),[]byte(password))
	if err != nil {
		return user,err
	}

	return user,nil


}

func (s *userService)IsCheckEmail (input CheckEmailInput) (bool,error){
	email := input.Email
	user,err := s.repository.FindByEmail(email)
	if err != nil {
		return false,err
	}

	if user.ID == 0 {
		return true,nil
	}

	return false,nil
}

func (s *userService)SaveAvatar (ID int,fileLocation string) (User,error){
	//dapatkan user berdasarkan id
	//update atrribute avatar
	//simpan perubahan
	user, err := s.repository.FindById(ID)
	if err != nil {
		return user,err
	}
	user.AvatarFileName = fileLocation
	updateUser,err := s.repository.UpdateUser(user)
	if err != nil {
		return updateUser,err
	}

	return updateUser,nil
}

func (s *userService)GetUserByID(ID int) (User,error){
	user,err := s.repository.FindById(ID)
	if err != nil {
		return user,err
	}

	if user.ID == 0 {
		return user,errors.New("No user found in that user id")
	}

	return user,nil

}