package todo_handler

import (
	"fmt"
	"net/http"
	logger "todo/tools"

	"github.com/gin-gonic/gin"
)

func (h *TodoHandler) GetTodos(c *gin.Context) {
	todos, err := h.Prisma.Todo.FindMany().Exec(c)

	if err != nil {
		logger.Logger.Error(fmt.Sprintf("Error fetching todos: %v\n", err));

		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Error fetching todos"},
		)
	}

	c.JSON(http.StatusOK, todos)
}