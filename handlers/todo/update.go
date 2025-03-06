package todo_handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"todo/models"
	"todo/prisma/db"
	logger "todo/tools"
)

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

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	todoString := fmt.Sprintf("%+v", todo);

	logger.Logger.Info("Updated record: " +  todoString);

	c.JSON(http.StatusOK, todo);
}
