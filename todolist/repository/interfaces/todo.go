package interfaces

import (
	"context"
	"time"

	"github.com/k0825/go-gin-ent-sample/models"
)

type TodoRepositoryInterface interface {
	FindById(context.Context, models.TodoId) (*models.Todo, error)
	Create(context.Context,
		models.TodoTitle,
		models.TodoDescription,
		models.TodoImage,
		[]models.TodoTag,
		time.Time,
		time.Time) (*models.Todo, error)
}
