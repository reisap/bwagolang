package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golangbwa/helper"
	"golangbwa/user"
	"net/http"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler{
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context){
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		var errors []string
		for _,e := range err.(validator.ValidationErrors){
			errors = append(errors,e.Error())
		}
		errorMessage := gin.H{"error":errors}
		response := helper.APIResponse("Register account Failed",http.StatusUnprocessableEntity,"error",errorMessage)
		c.JSON(http.StatusBadRequest,response)
		return
	}

	newUser,err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account Failed",http.StatusBadRequest,"error",nil)
		c.JSON(http.StatusBadRequest,response)
		return
	}
	formatter := user.FormatUser(newUser,"token")
	response := helper.APIResponse("Account has been registered",http.StatusOK,"success",formatter)

	c.JSON(http.StatusOK,response)

}