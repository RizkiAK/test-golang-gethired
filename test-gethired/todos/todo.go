package todos

import (
	"fmt"
	"net/http"
	"strconv"
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
	activityGroupID, _ := strconv.Atoi(c.Query("activity_group_id"))

	todolist, err := h.service.GetAll(activityGroupID)
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
		c.JSON(http.StatusNotFound, response)
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
		response := helper.APIResponse("title cannot be null", "Bad Request", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if input.ActivityGroupID == 0 {
		response := helper.APIResponse("activity_group_id cannot be null", "Bad Request", nil)
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

	todolist, _ := h.service.GetOne(inputID.ID)
	if todolist.ID == 0 {
		msg := fmt.Sprintf("Todo with ID %d Not Found", inputID.ID)
		response := helper.APIResponse(msg, "Not Found", nil)
		c.JSON(http.StatusNotFound, response)
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
		c.JSON(http.StatusNotFound, response)
		return
	}

	err = h.service.Delete(input.ID)
	if err != nil {
		response := helper.APIResponse("Error", "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success", "Success", struct{}{})
	c.JSON(http.StatusOK, response)

}
