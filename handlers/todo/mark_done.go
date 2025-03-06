package todo_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"todo/prisma/db"
)

func (h *TodoHandler) MarkDone(c *gin.Context) {
	id := c.Param("id")

    _, err := h.Prisma.Todo.FindUnique(
        db.Todo.ID.Equals(id),
    ).Update(
        db.Todo.Completed.Set(true),
    ).Exec(c)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.Status(http.StatusOK)
}