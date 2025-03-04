package models

type Todo struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
    Done  bool   `json:"done"`
}

type CreateTodoBody struct {
   Title string `json:"title"`
   Description string `json:"description"`
}

type UpdateTodoBody struct {
   Title string `json:"title"`
   Description string `json:"description"`
}