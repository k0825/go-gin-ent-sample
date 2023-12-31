package implements

import (
	"context"

	"github.com/cockroachdb/errors"
	repositoryinterfaces "github.com/k0825/go-gin-ent-sample/repository/interfaces"
	"github.com/k0825/go-gin-ent-sample/usecase/interfaces"
)

type TodoDeleteInteractor struct {
	todoRepository repositoryinterfaces.TodoRepositoryInterface
}

func NewTodoDeleteInteractor(tr repositoryinterfaces.TodoRepositoryInterface) *TodoDeleteInteractor {
	return &TodoDeleteInteractor{
		todoRepository: tr,
	}
}

func (tdi *TodoDeleteInteractor) Handle(ctx context.Context, request interfaces.TodoDeleteRequest) error {
	if tdi == nil {
		return errors.New("TodoDeleteInteractor is nil.")
	}

	_, err := tdi.todoRepository.RunInTx(ctx, func(ctx context.Context) (interface{}, error) {
		err := tdi.todoRepository.Delete(ctx, request.Id)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return nil, nil
	})

	if err != nil {
		return err
	}

	return nil
}
