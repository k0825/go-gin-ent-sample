package interfaces

import (
	"context"
	"time"

	"github.com/k0825/go-gin-ent-sample/models"
)

type TodoRepositoryInterface interface {
	FindById(context.Context, models.TodoId) (*models.Todo, error)
	FindAll(context.Context, int, int) ([]*models.Todo, error)
	Create(context.Context,
		models.TodoTitle,
		models.TodoDescription,
		models.TodoImage,
		[]models.TodoTag,
		time.Time,
		time.Time) (*models.Todo, error)
	// Update(context.Context,
	// 	models.TodoId,
	// 	*models.TodoTitle,
	// 	*models.TodoDescription,
	// 	*models.TodoImage,
	// 	*[]models.TodoTag,
	// 	*time.Time,
	// 	*time.Time) error
	Delete(context.Context, models.TodoId) error
}
