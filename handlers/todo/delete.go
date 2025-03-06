package todo_handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"todo/prisma/db"
	logger "todo/tools"
)

const (
	RecordNotFoundToDeleteError = "Record not found to delete"	
	RecordDeletionError = "Error deleting todo"
)

func (h *TodoHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	todo, err := h.Prisma.Todo.FindUnique(
		db.Todo.ID.Equals(id),
	).Delete().Exec(c)

	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			logger.Logger.Error(fmt.Sprintf(RecordNotFoundToDeleteError+": %v\n", err))

			c.JSON(http.StatusBadRequest, gin.H{"error": RecordNotFoundToDeleteError})
			return
		}
		
		logger.Logger.Error(fmt.Sprintf(RecordDeletionError + ": %v\n", err));

		c.JSON(http.StatusInternalServerError, gin.H{"error": RecordDeletionError})
		return
	}

	logger.LogWithObject("Deleted record", todo);

	c.Status(http.StatusNoContent)
}
