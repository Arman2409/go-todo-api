package todo_handler

import (
	"todo/prisma/db"
)

type TodoHandler struct {
	Prisma *db.PrismaClient
}