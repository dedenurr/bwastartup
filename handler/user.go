package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context)  {
	// tangkap input dari user
	// map input dari user ke struct RegisterInput
	// struct di atas kita pasing sebagai parameter service

	var input user.RegisterUserinput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}


		response := helper.APIResponse("Registered Account Failed", http.StatusUnprocessableEntity,"error",errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Registered Account Failed", http.StatusBadRequest,"error",nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser,"tokentokentokrn")

	response := helper.APIResponse("Account has been registered", http.StatusOK,"success",formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler)  Login(c *gin.Context){
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("login Failed", http.StatusUnprocessableEntity,"error",errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser,err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login Failed", http.StatusBadRequest,"error",errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loggedinUser,"tokentokentokrn")

	response := helper.APIResponse("Succesfuly Loggedin", http.StatusOK,"success",formatter)

	c.JSON(http.StatusOK, response)


}