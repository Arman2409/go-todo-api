package todo_handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"todo/models"
	"todo/prisma/db"
)


func (h *TodoHandler) CreateTodo(c *gin.Context) {

	var requestBody models.CreateTodoBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := h.Prisma.Todo.CreateOne(
		db.Todo.Title.Set(requestBody.Title),
		db.Todo.Description.Set(requestBody.Description),
	).Exec(c)

	if err != nil {
		fmt.Printf("Error creating todo: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating todo"})
	}

	c.JSON(http.StatusCreated, todo)
}