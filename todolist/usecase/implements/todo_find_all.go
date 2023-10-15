package implements

import (
	"context"

	"github.com/cockroachdb/errors"
	repositoryinterfaces "github.com/k0825/go-gin-ent-sample/repository/interfaces"
	"github.com/k0825/go-gin-ent-sample/usecase/interfaces"
)

type TodoFindAllInteractor struct {
	TodoRepository repositoryinterfaces.TodoRepositoryInterface
}

func NewTodoFindAllInteractor(tr repositoryinterfaces.TodoRepositoryInterface) *TodoFindAllInteractor {
	return &TodoFindAllInteractor{
		TodoRepository: tr,
	}
}

func (tfi *TodoFindAllInteractor) Handle(ctx context.Context, request interfaces.TodoFindAllRequest) (*interfaces.TodoFindAllResponse, error) {
	if tfi == nil {
		return nil, errors.New("TodoFindAllInteractor is nil.")
	}

	todos, err := tfi.TodoRepository.FindAll(ctx, request.Page, request.Number)

	if err != nil {
		return nil, err
	}

	response := interfaces.NewTodoFindAllResponse(todos)

	return response, nil
}
