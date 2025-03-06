package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"todo/handlers"
	"todo/prisma/db"
)

func main() {
    router := gin.Default()

    prismaClient := db.NewClient();

    err := prismaClient.Prisma.Connect()
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }
    
    defer prismaClient.Prisma.Disconnect() 
    
    todoHandler := handlers.TodoHandler{Prisma: prismaClient}

    router.GET("/todos", todoHandler.GetTodos)
    router.POST("/todo", todoHandler.CreateTodo)
    router.PATCH("/todo/:id", todoHandler.UpdateTodo)
    router.PATCH("/todo/:id/done",todoHandler.MarkDone)
    router.DELETE("/todo/:id",todoHandler.Delete)

    router.Run(":8080")
}