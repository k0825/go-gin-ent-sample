package implements

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/k0825/go-gin-ent-sample/domainerrors"
	"github.com/k0825/go-gin-ent-sample/models"
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
		return nil, errors.New("TodoCreateInteractor is nil.")
	}

	v, err := tci.todoRepository.RunInTx(ctx, func(ctx context.Context) (interface{}, error) {
		v, err := tci.todoRepository.Create(ctx, request.Title, request.Description, request.Image, request.Tags, request.StartsAt, request.EndsAt)

		if err != nil {
			return nil, errors.WithStack(err)
		}
		return v, nil
	})

	if err != nil {
		return nil, err
	}

	todo, ok := v.(*models.Todo)

	if !ok {
		intErr := domainerrors.NewInternalServerError("TodoCreateResponse value is incorrect")
		return nil, errors.WithStack(intErr)
	}

	response := interfaces.NewTodoCreateResponse(*todo)

	return response, nil
}
