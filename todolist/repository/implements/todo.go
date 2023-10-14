package implements

import (
	"context"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/k0825/go-gin-ent-sample/domainerrors"
	"github.com/k0825/go-gin-ent-sample/ent"
	"github.com/k0825/go-gin-ent-sample/ent/todo"
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
		return nil, errors.New("TodoRepositoryInterface pointer is nil")
	}

	todo, err := tr.client.Todo.Query().Where(todo.ID(todoId.Value())).WithTags().Only(ctx)

	if err != nil {
		nfErr := domainerrors.NewNotFoundError("Todo", todoId.String())
		wrapErr := errors.WithStack(nfErr)
		return nil, wrapErr
	}

	tags := todo.Edges.Tags

	if tags == nil {
		nfErr := domainerrors.NewNotFoundError("Tag", todoId.String())
		wrapErr := errors.WithStack(nfErr)
		return nil, wrapErr
	}

	dt, err := models.ConvertEntToTodo(todo, tags)

	if err != nil {
		return nil, err
	}

	return dt, nil
}

func (tr *TodoRepository) Create(ctx context.Context,
	title domain.TodoTitle,
	description domain.TodoDescription,
	image domain.TodoImage,
	tags []domain.TodoTag,
	startsAt time.Time,
	endsAt time.Time) (*domain.Todo, error) {

	if tr == nil {
		return nil, errors.New("TodoRepositoryInterface pointer is nil")
	}

	tx, err := tr.client.Tx(ctx)

	if err != nil {
		return nil, err
	}

	todoEnt, err := tx.Todo.Create().
		SetTitle(title.Value()).
		SetDescription(description.Value()).
		SetImage(image.Value()).
		SetStartsAt(startsAt).
		SetEndsAt(endsAt).
		Save(ctx)

	if err != nil {
		return nil, rollback(tx, err)
	}

	tagEnt, err := tx.Tag.MapCreateBulk(tags, func(c *ent.TagCreate, i int) {
		c.SetKeyword(tags[i].Value()).
			SetTodoID(todoEnt.ID)
	}).Save(ctx)

	if err != nil {
		return nil, rollback(tx, err)
	}

	tagKeywords := make([]string, len(tagEnt))
	for i, tag := range tagEnt {
		tagKeywords[i] = tag.Keyword
	}

	mtm := models.NewTodoModel(
		todoEnt.ID,
		todoEnt.Title,
		todoEnt.Description,
		*todoEnt.Image,
		tagKeywords,
		todoEnt.StartsAt,
		todoEnt.EndsAt,
		todoEnt.CreatedAt,
		todoEnt.UpdatedAt,
	)

	dt, err := mtm.ConvertToTodo()

	if err != nil {
		return nil, err
	}

	err = tx.Commit()

	return dt, err
}

func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = errors.Wrap(err, rerr.Error())
	}
	return err
}
