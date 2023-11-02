package implements

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/k0825/go-gin-ent-sample/domainerrors"
	"github.com/k0825/go-gin-ent-sample/models"
	repositoryinterfaces "github.com/k0825/go-gin-ent-sample/repository/interfaces"
	"github.com/k0825/go-gin-ent-sample/usecase/interfaces"
)

type TodoUpdateInteractor struct {
	todoRepository repositoryinterfaces.TodoRepositoryInterface
}

func NewTodoUpdateInteractor(tr repositoryinterfaces.TodoRepositoryInterface) *TodoUpdateInteractor {
	return &TodoUpdateInteractor{
		todoRepository: tr,
	}
}

func (tci *TodoUpdateInteractor) Handle(ctx context.Context, request interfaces.TodoUpdateRequest) (*interfaces.TodoUpdateResponse, error) {
	if tci == nil {
		return nil, errors.New("TodoUpdateInteractor is nil.")
	}

	v, err := tci.todoRepository.RunInTx(ctx, func(ctx context.Context) (interface{}, error) {
		_, err := tci.todoRepository.FindById(ctx, request.Id)

		if err != nil {
			return nil, err
		}

		v, err := tci.todoRepository.Update(ctx, request.Id, request.Title, request.Description, request.Image, request.Tags, request.StartsAt, request.EndsAt)

		if err != nil {
			return nil, err
		}
		return v, nil
	})

	if err != nil {
		return nil, err
	}

	todo, ok := v.(*models.Todo)

	if !ok {
		intErr := domainerrors.NewInternalServerError("TodoUpdateResponse value is incorrect")
		return nil, errors.WithStack(intErr)
	}

	response := interfaces.NewTodoUpdateResponse(*todo)

	return response, nil
}
