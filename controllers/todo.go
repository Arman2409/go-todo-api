package controllers

import (
    "github.com/gin-gonic/gin"
    "todo/handlers/todo"
	"todo/prisma/db"
)

type TodoController struct {
    TodoHandler *todo_handler.TodoHandler
}

func NewTodoController(client *db.PrismaClient) *TodoController {
    return &TodoController{
        TodoHandler: &todo_handler.TodoHandler{
            Prisma: client,
        },
    }
}

func (tc *TodoController) SetupRoutes(r *gin.Engine) {
    todosGroup := r.Group("/todos");

    {
        todosGroup.GET("", tc.TodoHandler.GetTodos)
        todosGroup.POST("", tc.TodoHandler.CreateTodo)
        todosGroup.PATCH("/:id", tc.TodoHandler.UpdateTodo)
        todosGroup.PATCH("/:id/done", tc.TodoHandler.MarkDone)
        todosGroup.DELETE("/:id", tc.TodoHandler.Delete)
    }
}

