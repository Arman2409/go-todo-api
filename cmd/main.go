package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"todo/controllers"
	"todo/prisma/db"
	logger "todo/tools"
)

func main() {
    logger.InitLogger()

    router := gin.Default()

    prismaClient := db.NewClient();

    defer prismaClient.Prisma.Disconnect() 

    err := prismaClient.Prisma.Connect()
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }
    
    todoController := controllers.NewTodoController(prismaClient)

    todoController.SetupRoutes(router);

    router.Run(":8080")
}