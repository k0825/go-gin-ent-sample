package implements

import (
	"context"

	"github.com/cockroachdb/errors"
	repositoryinterfaces "github.com/k0825/go-gin-ent-sample/repository/interfaces"
	"github.com/k0825/go-gin-ent-sample/usecase/interfaces"
)

type TodoExportInteractor struct {
	TodoRepository repositoryinterfaces.TodoRepositoryInterface
}

func NewTodoExportInteractor(tr repositoryinterfaces.TodoRepositoryInterface) *TodoExportInteractor {
	return &TodoExportInteractor{
		TodoRepository: tr,
	}
}

func (tei *TodoExportInteractor) Handle(ctx context.Context) (*interfaces.TodoExportResponse, error) {
	if tei == nil {
		return nil, errors.New("TodoExportInteractor is nil.")
	}

	todos, err := tei.TodoRepository.Export(ctx)

	if err != nil {
		return nil, err
	}

	response := interfaces.NewTodoExportResponse(todos)

	return response, nil
}
