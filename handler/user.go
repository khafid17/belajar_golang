package handler

import (
	"bwastartup/helper"
	"bwastartup/user"

	// "go/format"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserhandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatterValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Registered account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Registered account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// token, err := h.jwtService.GeneratorToken()
	formatter := user.FormatterUser(newUser, "tokentokentokentoken")

	response := helper.APIResponse("Account has been Registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}
