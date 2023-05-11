package handlers

import (
	"net/http"
	"strconv"
	"test/dto"
	tododto "test/dto/todo"
	"test/models"
	"test/repositories"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type handlerTodo struct {
	TodoRepository repositories.TodoRepository
}

func HandlerTodo(TodoRepository repositories.TodoRepository) *handlerTodo {
	return &handlerTodo{TodoRepository}
}

// get  ALLusers
// @Summary CRUD  GET ALLUSER
// @Produce json
// @Router /users [get]
func (h *handlerTodo) GetTodos(c *gin.Context) {
	activityGroupID, err := strconv.Atoi(c.Query("activity_group_id"))
	var todos []models.Todo
	if err != nil {
		todos, _ = h.TodoRepository.GetAllTodo()
		response := dto.SuccessResult{Status: "Success", Message: "Success", Data: todos}
		c.JSON(http.StatusOK, response)
		return
	}

	todos, _ = h.TodoRepository.GetTodosByActivityGroupID(activityGroupID)

	response := dto.SuccessResult{Status: "Success", Message: "Success", Data: todos}
	c.JSON(http.StatusOK, response)
}

func (h *handlerTodo) CreateTodo(c *gin.Context) {
	var request tododto.TodoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//validator
	validate := validator.New()

	if err := validate.Struct(request); err != nil {
		response := dto.ErrorResult{Status: "Bad request", Message: "Title cannot be null"}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	todoData := models.Todo{
		Title:           request.Title,
		ActivityGroupID: request.ActivityGroupID,
		IsActive:        request.IsActive,
		Priority:        "very-high",
	}

	todo, err := h.TodoRepository.TodoCreate(todoData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.SuccessResult{Status: "Success", Message: "Success", Data: todo}
	c.JSON(http.StatusCreated, response)
}

// getuser by id
// @Summary get a user item by ID
// @ID get-user-by-id
// @Produce json
// @Param id path string true "user ID"
// @Router /user/{id} [get]
func (h *handlerTodo) GetOneTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := h.TodoRepository.GetTodo(id)
	if err != nil {
		response := dto.ErrorResult{Status: "Not found", Message: "Activity with ID " + strconv.Itoa(id) + " not found"}
		c.JSON(http.StatusNotFound, response)
		return
	}
	response := dto.SuccessResult{Status: "Success", Message: "Success", Data: todo}
	c.JSON(http.StatusOK, response)
}

// @Summary delete a user by ID
// @ID delete-user-by-id
// @Produce json
// @Param id path string true "user ID"
// @Router /user/{id} [delete]
func (h *handlerTodo) TodoDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	todo, err := h.TodoRepository.GetTodo(id)
	if err != nil {
		response := dto.ErrorResult{Status: "Not found", Message: "Todo with ID " + strconv.Itoa(id) + " not found"}
		c.JSON(http.StatusNotFound, response)
		return
	}

	data, _ := h.TodoRepository.DeleteTodo(todo)

	deleteResponse := tododto.DeleteResponse{
		TodoID: data.TodoID,
	}

	response := dto.SuccessResult{Status: "Success", Message: "Success", Data: deleteResponse}
	c.JSON(http.StatusOK, response)
}

func (h *handlerTodo) TodoUpdate(c *gin.Context) {
	var request tododto.TodoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//validator
	validate := validator.New()

	if err := validate.Struct(request); err != nil {
		response := dto.ErrorResult{Status: "Bad request", Message: "Title cannot be null"}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := h.TodoRepository.GetTodo(id)
	if err != nil {
		response := dto.ErrorResult{Status: "Not found", Message: "Activity with ID " + strconv.Itoa(id) + " not found"}
		c.JSON(http.StatusNotFound, response)
		return
	}

	if request.Title != "" {
		todo.Title = request.Title
	}

	if request.Priority != "" {
		todo.Priority = request.Priority
	}

	if request.IsActive != false {
		todo.IsActive = request.IsActive
	}

	todoUpdate, err := h.TodoRepository.UpdateTodo(todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.SuccessResult{Status: "Success", Message: "Success", Data: todoUpdate}
	c.JSON(http.StatusOK, response)
}
