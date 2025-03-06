package todo_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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