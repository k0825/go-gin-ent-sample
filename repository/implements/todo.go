package implements

import (
	"context"
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/k0825/go-gin-ent-sample/domainerrors"
	"github.com/k0825/go-gin-ent-sample/ent"
	"github.com/k0825/go-gin-ent-sample/ent/tag"
	domain "github.com/k0825/go-gin-ent-sample/models"
	"github.com/k0825/go-gin-ent-sample/repository/models"
)

type TodoRepository struct {
	client *ent.Client
}

func NewTodoRepository(client *ent.Client) (*TodoRepository, error) {
	return &TodoRepository{client}, nil
}

func (tr *TodoRepository) FindById(ctx context.Context, todoId domain.TodoId) (*domain.Todo, error) {
	if tr == nil {
		return nil, fmt.Errorf("TodoRepositoryInterface pointer is nil")
	}

	todo, err := tr.client.Todo.Get(ctx, todoId.Value())

	if err != nil {
		nfErr := domainerrors.NewNotFoundError("Todo", todoId.String())
		wrapErr := errors.Wrap(err, nfErr.Error())
		return nil, wrapErr
	}

	tags, err := tr.client.Tag.Query().Where(tag.TodoID(todoId.Value())).Select(tag.FieldKeyword).Strings(ctx)

	if err != nil {
		nfErr := domainerrors.NewNotFoundError("Tag", todoId.String())
		wrapErr := errors.Wrap(err, nfErr.Error())
		return nil, wrapErr
	}

	mtm := models.NewTodoModel(
		todo.ID,
		todo.Title,
		todo.Description,
		*todo.Image,
		tags,
		todo.StartsAt,
		todo.EndsAt,
		todo.CreatedAt,
		todo.UpdatedAt,
	)

	dt, err := mtm.ConvertToTodo()

	if err != nil {
		return nil, err
	}

	return dt, nil
}
