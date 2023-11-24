package interfaces

import (
	"context"

	"github.com/k0825/go-gin-ent-sample/models"
)

type TodoExportResponse struct {
	Todos []*models.Todo
}

func NewTodoExportResponse(todos []*models.Todo) *TodoExportResponse {
	return &TodoExportResponse{
		Todos: todos,
	}
}

type TodoExportUseCaseInterface interface {
	Handle(context.Context) (*TodoExportResponse, error)
}
