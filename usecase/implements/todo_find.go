package implements

import (
	"context"
	"fmt"

	repositoryinterfaces "github.com/k0825/go-gin-ent-sample/repository/interfaces"
	"github.com/k0825/go-gin-ent-sample/usecase/interfaces"
)

type TodoFindByIdInteractor struct {
	todoRepository repositoryinterfaces.TodoRepositoryInterface
}

func NewTodoFindByIdInteractor(tr repositoryinterfaces.TodoRepositoryInterface) *TodoFindByIdInteractor {
	return &TodoFindByIdInteractor{
		todoRepository: tr,
	}
}

func (tfi *TodoFindByIdInteractor) Handle(ctx context.Context, request interfaces.TodoFindRequest) (*interfaces.TodoFindResponse, error) {
	if tfi == nil {
		return nil, fmt.Errorf("TodoFindByIdInteractor is nil.")
	}

	todo, err := tfi.todoRepository.FindById(ctx, request.Id)

	if err != nil {
		return nil, err
	}

	response := interfaces.NewTodoFindResponse(*todo)

	return response, nil
}
