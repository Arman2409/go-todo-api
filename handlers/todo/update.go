package todo_handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"todo/models"
	"todo/prisma/db"
	logger "todo/tools"
)

const (
	FailedToUpdateError = "Failed to update the record"
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
		if errors.Is(err, db.ErrNotFound) {
			logger.Logger.Error(fmt.Sprintf(RecordNotFoundForUpdateError+": %v\n", err))

			c.JSON(http.StatusBadRequest, gin.H{"error": RecordNotFoundForUpdateError})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": FailedToUpdateError})
		return
	}

	logger.LogWithObject("Updated record", todo)

	c.JSON(http.StatusOK, todo)
}
