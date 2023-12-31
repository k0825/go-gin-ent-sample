package main

import (
	"github.com/gin-gonic/gin"
	"github.com/k0825/go-gin-ent-sample/container"
	_ "github.com/lib/pq"
)

func setupRouter() *gin.Engine {

	c, err := container.Init()

	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/todos", c.TodoController.GetAllTodo)
	r.GET("/todos/:id", c.TodoController.GetTodo)
	r.DELETE("/todos/:id", c.TodoController.DeleteTodo)
	r.POST("/todos", c.TodoController.PostTodo)
	r.PUT("/todos/:id", c.TodoController.PutTodo)
	r.GET("/todos/search", c.TodoController.SearchTodo)

	r.GET("/export", c.TodoController.ExportTodo)

	return r
}

func main() {
	r := setupRouter()
	if err := r.Run(); err != nil {
		panic(err)
	}
}
