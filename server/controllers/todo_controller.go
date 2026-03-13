package controllers

import (
	"net/http"

	"github.com/sujin/todo-app/database"
	"github.com/sujin/todo-app/models"
	"github.com/gin-gonic/gin"
)

// Create Todo

// CreateTodo godoc
// @Summary Create a new todo
// @Description Create a new todo for logged-in user
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body models.TodoReq true "Todo payload"
// @Success 201 {object} models.TodoReq
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /todos [post]
// @Security BearerAuth
func CreateTodo(c *gin.Context) {
	var todo models.TodoReq
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("user_id")
	newTodo := models.Todo{
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   false,
		UserID:      userID,
	}

	if err := database.DB.Create(&newTodo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

// Get all todos for logged-in user

// GetTodos godoc
// @Summary List all todos 
// @Description Get all todos for the logged-in user
// @Tags todos
// @Accept json
// @Produce json
// @Success 200 {object} []models.TodoResponse
// @Failure 401 {object} map[string]string
// @Router /todos [get]
// @Security BearerAuth
func GetTodos(c *gin.Context) {
	userID := c.GetUint("user_id")
	var todos []models.Todo

	if err := database.DB.Where("user_id = ?", userID).Find(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
		return
	}

	var res = make([]models.TodoResponse, 0, len(todos))
	for _, todo := range todos {
		res = append(res, models.TodoResponse{
			ID:        todo.ID,
			Title:     todo.Title,
			Completed: todo.Completed,
			CreatedAt: todo.CreatedAt,
			UpdatedAt: todo.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, res)
}

// Get single todo by ID

// GetTodo godoc
// @Summary Get a single todo
// @Description Get a todo by its ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} models.TodoResponse
// @Failure 400 {object} map[string]string
// @Router /todos/{id} [get]
// @Security BearerAuth
func GetTodoByID(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	 res := models.TodoResponse{
		ID: todo.ID,
		Title: todo.Title,
		Completed: todo.Completed,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	 }

	 
	c.JSON(http.StatusOK, res)
}

// Update todo

// UpdateTodo godoc
// @Summary Update a todo
// @Description Update a todo by its ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Param todo body models.TodoReq true "Todo payload"
// @Success 200 {object} models.TodoReq
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /todos/{id} [put]
// @Security BearerAuth
func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var input models.TodoReq
	var todo models.Todo

	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo.Title = input.Title
	todo.Description = input.Description
	todo.Completed = input.Completed

	database.DB.Save(&todo)
	c.JSON(http.StatusOK, gin.H{"message": "Todo updated", "todo": input})
}

// Delete todo

// DeleteTodo godoc
// @Summary Delete a todo
// @Description Delete a todo by its ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Success 204 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /todos/{id} [delete]
// @Security BearerAuth
func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	database.DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
