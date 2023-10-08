package main

import (
	"github.com/gin-gonic/gin"
	"github.com/k0825/go-gin-ent-sample/config"
	"github.com/k0825/go-gin-ent-sample/controller"
	"github.com/k0825/go-gin-ent-sample/datasource"
	repository "github.com/k0825/go-gin-ent-sample/repository/implements"
	usecase "github.com/k0825/go-gin-ent-sample/usecase/implements"
	_ "github.com/lib/pq"
)

func setupRouter() *gin.Engine {

	conf := config.NewConfig()
	client, err := datasource.NewConnection(conf)

	if err != nil {
		panic(err)
	}

	adRepository, err := repository.NewTodoRepository(client)

	if err != nil {
		panic(err)
	}

	adFindByIdUsecase := usecase.NewTodoFindByIdInteractor(adRepository)
	ctrl := controller.NewTodoController(adFindByIdUsecase)

	r := gin.Default()
	r.GET("/todos/:id", func(ctx *gin.Context) {
		ctrl.GetTodo(ctx)
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run()
}
