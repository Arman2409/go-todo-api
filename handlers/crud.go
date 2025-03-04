package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"todo/models"
	"todo/prisma/db"
)

type TodoHandler struct {
	Prisma *db.PrismaClient
}

func (h *TodoHandler) GetTodos(c *gin.Context) {
	todos, err := h.Prisma.Todo.FindMany().Exec(c)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Error fetching todos"},
		)
	}

	c.JSON(http.StatusOK, todos)
}

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

func (h *TodoHandler) UpdateTodo(c *gin.Context) {

	id := c.Param("id")

	var requestBody models.UpdateTodoBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read the request body"})
		return
	}

	updateArgs := []db.TodoSetParam{}

	if requestBody.Title != "" { 
		updateArgs = append(updateArgs, db.Todo.Title.Set(requestBody.Title))
	}

	if requestBody.Description != "" { 
		updateArgs = append(updateArgs, db.Todo.Description.Set(requestBody.Description))
	}

	if len(updateArgs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No fields to update"})
		return
	}

	todo, err := h.Prisma.Todo.FindUnique(
		db.Todo.ID.Equals(id),
	).Update(
		updateArgs...,
	).Exec(c)

	fmt.Println(todo)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo);
}
