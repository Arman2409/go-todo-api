package todo_handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"todo/prisma/db"
	logger "todo/tools"
)

func (h *TodoHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	todo, err := h.Prisma.Todo.FindUnique(
		db.Todo.ID.Equals(id),
	).Delete().Exec(c)

	if err != nil {
		logger.Logger.Error(fmt.Sprintf("Error deleting todo: %v\n", err));

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting todo"})
		return
	}

	todoString := fmt.Sprintf("%+v", todo)

	logger.Logger.Info("Deleted record: " +  todoString);

	c.Status(http.StatusNoContent)
}
