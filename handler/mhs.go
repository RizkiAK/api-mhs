package handler

import (
	"api-data-mhs/helper"
	"api-data-mhs/mhs"
	"api-data-mhs/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type mhsHandler struct {
	service mhs.ServiceInterface
}

func NewMhsHandler(service mhs.ServiceInterface) *mhsHandler {
	return &mhsHandler{service}
}

func (h *mhsHandler) Create(c *gin.Context) {
	var input mhs.Mahasiswa

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationErr(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Create data failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userID := c.MustGet("currentUser").(user.User)

	input.ID = userID.Nim

	err = h.service.Create(input)
	if err != nil {
		response := helper.APIResponse("Create data failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Create data success", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *mhsHandler) Update(c *gin.Context) {
	var input mhs.InputMhsDetail

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Update data failed", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var mhs mhs.InputMhs
	err = c.ShouldBindJSON(&mhs)
	if err != nil {
		errors := helper.FormatValidationErr(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Update data failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user := c.MustGet("currentUser").(user.User)

	err = h.service.Update(input, mhs, user.Nim)
	if err != nil {
		response := helper.APIResponse("Update data failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Update data success", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *mhsHandler) Delete(c *gin.Context) {
	var input mhs.InputMhsDetail

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Delete data failed", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	user := c.MustGet("currentUser").(user.User)

	h.service.Delete(input, user.Nim)

	response := helper.APIResponse("Delete data success", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *mhsHandler) FindAll(c *gin.Context) {
	data, err := h.service.FindAll()
	if err != nil {
		response := helper.APIResponse("Get all data failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := mhs.FormatterArray(data)
	response := helper.APIResponse("Get all data success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *mhsHandler) FindByNim(c *gin.Context) {
	var input mhs.InputMhsDetail

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Get data failed", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user := c.MustGet("currentUser").(user.User)

	data, err := h.service.FindByNim(input, user.Nim)
	if err != nil {
		response := helper.APIResponse("Get data failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Get data success", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
