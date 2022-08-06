package handler

import (
	"final-project-sanbercode/helper"
	"final-project-sanbercode/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse(
			"Register User Failed",
			http.StatusUnprocessableEntity,
			"Error",
			errorMessage,
		)
		c.JSON(
			http.StatusUnprocessableEntity,
			response,
		)
		return
	}

	dataUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse(
			"Register User Failed",
			http.StatusBadRequest,
			"Error",
			nil,
		)
		c.JSON(
			http.StatusBadRequest,
			response,
		)
		return
	}

	formatter := user.FormatUser(dataUser, "TOKENSEMENTARA")

	response := helper.APIResponse(
		"Register User Success",
		http.StatusOK,
		"Success",
		formatter,
	)

	c.JSON(http.StatusOK, response)

}
