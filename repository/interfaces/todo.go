package interfaces

import (
	"context"

	"github.com/k0825/go-gin-ent-sample/models"
)

type TodoRepositoryInterface interface {
	FindById(context.Context, models.TodoId) (*models.Todo, error)
}
