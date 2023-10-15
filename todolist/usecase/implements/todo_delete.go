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

	err := tdi.todoRepository.Delete(ctx, request.Id)

	if err != nil {
		return err
	}

	return nil
}
