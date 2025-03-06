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
	MarkingDoneError    = "Error marking todo as done"
)

func (h *TodoHandler) MarkDone(c *gin.Context) {
	id := c.Param("id")

	todo, err := h.Prisma.Todo.FindUnique(
		db.Todo.ID.Equals(id),
	).Update(
		db.Todo.Completed.Set(true),
	).Exec(c)

	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			logger.Logger.Error(fmt.Sprintf(RecordNotFoundForUpdateError + ": %v\n", err))

			c.JSON(http.StatusBadRequest, gin.H{"error": RecordNotFoundForUpdateError})
			return;
		}

		logger.Logger.Error(fmt.Sprintf(MarkingDoneError + ": %v\n", err))

		c.JSON(http.StatusInternalServerError, gin.H{"error": MarkingDoneError})
		return
	}

	logger.LogWithObject("Marked done record", todo)

	c.Status(http.StatusOK)
}
