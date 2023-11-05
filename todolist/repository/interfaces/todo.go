package interfaces

import (
	"context"
	"time"

	"github.com/k0825/go-gin-ent-sample/models"
)

type TodoRepositoryInterface interface {
	FindById(context.Context, models.TodoId) (*models.Todo, error)
	FindByTitle(context.Context, string, int, int) ([]*models.Todo, *models.PaginationMeta, error)
	FindByDescription(context.Context, string, int, int) ([]*models.Todo, *models.PaginationMeta, error)
	FindByImage(context.Context, string, int, int) ([]*models.Todo, *models.PaginationMeta, error)
	FindByTag(context.Context, string, int, int) ([]*models.Todo, *models.PaginationMeta, error)
	FindAll(context.Context, int, int) ([]*models.Todo, *models.PaginationMeta, error)
	Create(context.Context,
		models.TodoTitle,
		models.TodoDescription,
		models.TodoImage,
		[]models.TodoTag,
		time.Time,
		time.Time) (*models.Todo, error)
	Update(context.Context,
		models.TodoId,
		models.TodoTitle,
		models.TodoDescription,
		models.TodoImage,
		[]models.TodoTag,
		time.Time,
		time.Time) (*models.Todo, error)
	Delete(context.Context, models.TodoId) error
	RunInTx(context.Context, func(context.Context) (any, error)) (any, error)
}
