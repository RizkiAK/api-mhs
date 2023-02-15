package handler

import (
	"api-data-mhs/auth"
	"api-data-mhs/helper"
	"api-data-mhs/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service user.ServiceInterface
	auth    auth.Service
}

func NewUserHandler(service user.ServiceInterface, auth auth.Service) *userHandler {
	return &userHandler{service, auth}
}

func (h *userHandler) Register(c *gin.Context) {
	var input user.User

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationErr(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = h.service.Register(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.UserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationErr(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := h.service.Login(input)
	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.auth.GenerateToken(user.Nim)
	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Login Success", http.StatusOK, "success", token)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) ForgotPassword(c *gin.Context) {
	var nim user.UserUri

	err := c.ShouldBindUri(&nim)
	if err != nil {
		response := helper.APIResponse("Reset password failed", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var input user.UserInput

	err = c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationErr(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Reset password failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = h.service.ForgotPassword(nim, input)
	if err != nil {
		response := helper.APIResponse("Reset password failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Reset password Success", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
