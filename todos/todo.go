package todos

import (
	"fmt"
	"net/http"
	"test-gethired/helper"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) GetAll(c *gin.Context) {
	todolist, err := h.service.GetAll()
	if err != nil {
		response := helper.APIResponse("Error", "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success", "Success", todolist)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) GetOne(c *gin.Context) {
	var input GetTodosDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Error", "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	todolist, err := h.service.GetOne(input.ID)
	if err != nil {
		response := helper.APIResponse("Error", "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if todolist.ID == 0 {
		msg := fmt.Sprintf("Todo with ID %d Not Found", input.ID)
		response := helper.APIResponse(msg, "Not Found", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success", "Success", todolist)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) Created(c *gin.Context) {
	var input InputTodos

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse("Error", "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if input.Title == "" {
		response := helper.APIResponse("Title cannot be null", "Bad Request", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	todolist, err := h.service.Created(input)
	if err != nil {
		response := helper.APIResponse("Error", "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success", "Success", todolist)
	c.JSON(http.StatusCreated, response)
}

func (h *Handler) Updated(c *gin.Context) {
	var inputID GetTodosDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Error", "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData InputTodosUpdate

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		response := helper.APIResponse("Error", "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if inputData.Title == "" {
		response := helper.APIResponse("Title cannot be null", "Bad Request", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	todolist, _ := h.service.GetOne(inputID.ID)
	if todolist.ID == 0 {
		msg := fmt.Sprintf("Todo with ID %d Not Found", inputID.ID)
		response := helper.APIResponse(msg, "Not Found", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedTodolist, err := h.service.Updated(inputID.ID, inputData)
	if err != nil {
		response := helper.APIResponse("Error", "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success", "Success", updatedTodolist)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) Delete(c *gin.Context) {
	var input GetTodosDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Error", "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	todolist, err := h.service.GetOne(input.ID)
	if err != nil {
		response := helper.APIResponse("Error", "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if todolist.ID == 0 {
		msg := fmt.Sprintf("Todo with ID %d Not Found", input.ID)
		response := helper.APIResponse(msg, "Not Found", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = h.service.Delete(input.ID)
	if err != nil {
		response := helper.APIResponse("Error", "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success", "Success", nil)
	c.JSON(http.StatusOK, response)

}
