package todo_handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"todo/prisma/db"
	logger "todo/tools"
)

func (h *TodoHandler) MarkDone(c *gin.Context) {
	id := c.Param("id")

	todo, err := h.Prisma.Todo.FindUnique(
		db.Todo.ID.Equals(id),
	).Update(
		db.Todo.Completed.Set(true),
	).Exec(c)

	if err != nil {
		logger.Logger.Error(fmt.Sprintf("Error marking todo as done: %v\n", err))

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error marking todo as done"})
		return
	}

	todoString := fmt.Sprintf("%+v", todo);

	logger.Logger.Info("Marked done record: " +  todoString);

	c.Status(http.StatusOK)
}
