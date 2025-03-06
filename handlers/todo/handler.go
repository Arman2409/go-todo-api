package todo_handler

import (
	"todo/prisma/db"
)

const (
	RecordNotFoundForUpdateError = "Record not found to update"	
)

type TodoHandler struct {
	Prisma *db.PrismaClient
}