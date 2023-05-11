package routes

import (
	"test/handlers"
	"test/pkg/mysql"
	"test/repositories"

	"github.com/gin-gonic/gin"
)

func TodoRoutes(r *gin.RouterGroup) {
	todoRepository := repositories.RepositoryTodo(mysql.DB)
	h := handlers.HandlerTodo(todoRepository)
	r.GET("/todo-items", h.GetTodos)
	r.POST("/todo-items", h.CreateTodo)
	r.GET("/todo-items/:id", h.GetOneTodo)
	r.DELETE("/todo-items/:id", h.TodoDelete)
	r.PATCH("/todo-items/:id", h.TodoUpdate)
	// r.HandleFunc("/user/update/{id}", h.UpdateUser).Methods("PATCH")
	// r.HandleFunc("/user/delete/{id}", h.Delete).Methods("DELETE")

}
