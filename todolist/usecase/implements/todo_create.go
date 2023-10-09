package implements

import (
	"context"
	"fmt"

	repositoryinterfaces "github.com/k0825/go-gin-ent-sample/repository/interfaces"
	"github.com/k0825/go-gin-ent-sample/usecase/interfaces"
)

type TodoCreateInteractor struct {
	todoRepository repositoryinterfaces.TodoRepositoryInterface
}

func NewTodoCreateInteractor(tr repositoryinterfaces.TodoRepositoryInterface) *TodoCreateInteractor {
	return &TodoCreateInteractor{
		todoRepository: tr,
	}
}

func (tci *TodoCreateInteractor) Handle(ctx context.Context, request interfaces.TodoCreateRequest) (*interfaces.TodoCreateResponse, error) {
	if tci == nil {
		return nil, fmt.Errorf("TodoCreateInteractor is nil.")
	}

	todo, err := tci.todoRepository.Create(ctx, request.Title, request.Description, request.Image, request.Tags, request.StartsAt, request.EndsAt)

	if err != nil {
		return nil, err
	}

	response := interfaces.NewTodoCreateResponse(*todo)

	return response, nil
}
