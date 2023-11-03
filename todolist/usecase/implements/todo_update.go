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

func (tui *TodoUpdateInteractor) Handle(ctx context.Context, request interfaces.TodoUpdateRequest) (*interfaces.TodoUpdateResponse, error) {
	if tui == nil {
		return nil, errors.New("TodoUpdateInteractor is nil.")
	}
	v, err := tui.todoRepository.RunInTx(ctx, func(ctx context.Context) (any, error) {
		_, err := tui.todoRepository.FindById(ctx, request.Id)

		if err != nil {
			return nil, err
		}

		v, err := tui.todoRepository.Update(ctx, request.Id, request.Title, request.Description, request.Image, request.Tags, request.StartsAt, request.EndsAt)

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
