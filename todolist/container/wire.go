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

var configSet = wire.NewSet(config.NewConfig)
var datasourceSet = wire.NewSet(datasource.NewRDBConnection)
var repositorySet = wire.NewSet(repository.NewTodoRepository)
var usecaseSet = wire.NewSet(usecase.NewTodoFindByIdInteractor, usecase.NewTodoFindAllInteractor, usecase.NewTodoCreateInteractor, usecase.NewTodoDeleteInteractor, usecase.NewTodoUpdateInteractor)
var controllerSet = wire.NewSet(controller.NewTodoController)

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
		configSet,
		datasourceSet,
		repositorySet,
		usecaseSet,
		controllerSet,
		newContainer,
		wire.Bind(new(datasource.RDBConnectionInterface), new(*datasource.RDBConnection)),
		wire.Bind(new(repositoryinterfaces.TodoRepositoryInterface), new(*repository.TodoRepository)),
		wire.Bind(new(usecaseinterfaces.TodoFindUseCaseInterface), new(*usecase.TodoFindByIdInteractor)),
		wire.Bind(new(usecaseinterfaces.TodoFindAllUseCaseInterface), new(*usecase.TodoFindAllInteractor)),
		wire.Bind(new(usecaseinterfaces.TodoCreateUseCaseInterface), new(*usecase.TodoCreateInteractor)),
		wire.Bind(new(usecaseinterfaces.TodoDeleteUseCaseInterface), new(*usecase.TodoDeleteInteractor)),
		wire.Bind(new(usecaseinterfaces.TodoUpdateUseCaseInterface), new(*usecase.TodoUpdateInteractor)),
		wire.Bind(new(controller.TodoControllerInterface), new(*controller.TodoController)),
	)
	return nil, nil
}
