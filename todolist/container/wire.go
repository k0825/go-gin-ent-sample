//go:build wireinject
// +build wireinject

package container

import (
	"github.com/google/wire"
	"github.com/k0825/go-gin-ent-sample/config"
	"github.com/k0825/go-gin-ent-sample/controller"
	"github.com/k0825/go-gin-ent-sample/datasource"
	repository "github.com/k0825/go-gin-ent-sample/repository/implements"
	repositoryinterfaces "github.com/k0825/go-gin-ent-sample/repository/interfaces"
	usecase "github.com/k0825/go-gin-ent-sample/usecase/implements"
	usecaseinterfaces "github.com/k0825/go-gin-ent-sample/usecase/interfaces"
)

var repositorySet = wire.NewSet(repository.NewTodoRepository)
var usecaseSet = wire.NewSet(usecase.NewTodoFindByIdInteractor, usecase.NewTodoCreateInteractor)
var controllerSet = wire.NewSet(controller.NewTodoController)
var configSet = wire.NewSet(config.NewConfig)

type Container struct {
	TodoController controller.TodoControllerInterface
}

func newContainer(controller controller.TodoControllerInterface) *Container {
	return &Container{
		TodoController: controller,
	}
}

func Init() (*Container, error) {
	wire.Build(
		datasource.NewConnection,
		configSet,
		repositorySet,
		usecaseSet,
		controllerSet,
		newContainer,
		wire.Bind(new(repositoryinterfaces.TodoRepositoryInterface), new(*repository.TodoRepository)),
		wire.Bind(new(usecaseinterfaces.TodoFindUseCaseInterface), new(*usecase.TodoFindByIdInteractor)),
		wire.Bind(new(usecaseinterfaces.TodoCreateUseCaseInterface), new(*usecase.TodoCreateInteractor)),
		wire.Bind(new(controller.TodoControllerInterface), new(*controller.TodoController)),
	)
	return nil, nil
}
