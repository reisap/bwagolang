package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
		formatError := helper.FormatValidationError(err)
		errorMessage := gin.H{"error":formatError}
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

func (h *userHandler) Login(c *gin.Context){
	//user input email & password
	//input ditangkap handler
	//mapping user ke service
	//service -> repository

	var input user.LoginInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		formatError := helper.FormatValidationError(err)
		errorMessage := gin.H{"error":formatError}
		response := helper.APIResponse("Login account Failed",http.StatusUnprocessableEntity,"error",errorMessage)
		c.JSON(http.StatusBadRequest,response)
		return
	}

	logginUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"erro":err.Error()}
		response := helper.APIResponse("Login account Failed",http.StatusUnprocessableEntity,"error",errorMessage)
		c.JSON(http.StatusBadRequest,response)
		return
	}

	formatter := user.FormatUser(logginUser,"token")
	response := helper.APIResponse("Success Login",http.StatusOK,"success",formatter)

	c.JSON(http.StatusOK,response)


}

func (h *userHandler) CheckEmailAvailablity(c *gin.Context){
	var input user.CheckEmailInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		formatError := helper.FormatValidationError(err)
		errorMessage := gin.H{"error":formatError}
		response := helper.APIResponse("Email account Failed",http.StatusUnprocessableEntity,"error",errorMessage)
		c.JSON(http.StatusBadRequest,response)
		return
	}
	isEmailAvailable, err := h.userService.IsCheckEmail(input)
	if err != nil {
		errorMessage := gin.H{"error":err.Error()}
		response := helper.APIResponse("Email account Failed",http.StatusUnprocessableEntity,"error",errorMessage)
		c.JSON(http.StatusBadRequest,response)
		return
	}

	metaMessage := "Email has been registered"
	if isEmailAvailable {
		metaMessage = "Email is availabe"
	}
	data := gin.H{"isAvailable" : isEmailAvailable}
	response := helper.APIResponse(metaMessage,http.StatusOK,"success",data)

	c.JSON(http.StatusOK,response)


}

func (h *userHandler)UploadAvatars (c *gin.Context){
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded":false}
		response := helper.APIResponse("Avatar uploaded Failed",http.StatusBadRequest,"error",data)
		c.JSON(http.StatusBadRequest,response)
	}
	userId := 1
	//path := "images/" + file.Filename
	path := fmt.Sprintf("images/%d-%s",userId,file.Filename)
	err = c.SaveUploadedFile(file,path)
	if err != nil {
		errorMessage := gin.H{"error":err.Error()}
		response := helper.APIResponse("Avatar uploaded Failed",http.StatusBadRequest,"error",errorMessage)
		c.JSON(http.StatusBadRequest,response)
		return
	}


	_,err = h.userService.SaveAvatar(userId,path)
	if err != nil {
		errorMessage := gin.H{"error":err.Error()}
		response := helper.APIResponse("Avatar uploaded Failed",http.StatusUnprocessableEntity,"error",errorMessage)
		c.JSON(http.StatusBadRequest,response)
		return
	}

	data := gin.H{"is_uploaded" : true}
	response := helper.APIResponse("Avatar user success uploaded",http.StatusOK,"success",data)

	c.JSON(http.StatusOK,response)


}