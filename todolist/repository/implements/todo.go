package implements

import (
	"context"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/k0825/go-gin-ent-sample/datasource"
	"github.com/k0825/go-gin-ent-sample/domainerrors"
	"github.com/k0825/go-gin-ent-sample/ent"
	"github.com/k0825/go-gin-ent-sample/ent/tag"
	"github.com/k0825/go-gin-ent-sample/ent/todo"
	domain "github.com/k0825/go-gin-ent-sample/models"
	"github.com/k0825/go-gin-ent-sample/repository/models"
)

type TodoRepository struct {
	conn datasource.RDBConnectionInterface
}

func NewTodoRepository(conn datasource.RDBConnectionInterface) (*TodoRepository, error) {
	return &TodoRepository{conn}, nil
}

func (tr *TodoRepository) FindById(ctx context.Context, todoId domain.TodoId) (*domain.Todo, error) {
	if tr == nil {
		return nil, errors.New("TodoRepositoryInterface pointer is nil")
	}

	client := tr.conn.GetTx(ctx)
	if client == nil {
		client = tr.conn.GetClient()
	}

	if client == nil {
		intErr := domainerrors.NewInternalServerError("client or transaction is not found")
		return nil, errors.WithStack(intErr)
	}

	todo, err := client.Todo.Query().Where(todo.ID(todoId.Value())).WithTags().Only(ctx)

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
		return nil, errors.WithStack(err)
	}

	return dt, nil
}

func (tr *TodoRepository) FindByTitle(ctx context.Context, title string, start int, take int) ([]*domain.Todo, *domain.PaginationMeta, error) {
	if tr == nil {
		return nil, nil, errors.New("TodoRepositoryInterface pointer is nil")
	}

	client := tr.conn.GetTx(ctx)
	if client == nil {
		client = tr.conn.GetClient()
	}

	if client == nil {
		intErr := domainerrors.NewInternalServerError("client or transaction is not found")
		return nil, nil, errors.WithStack(intErr)
	}

	todos, err := client.Todo.Query().Where(todo.TitleContains(title)).WithTags().Offset(start).Limit(take).All(ctx)

	if err != nil && !ent.IsNotFound(err) {
		return nil, nil, errors.WithStack(err)
	}

	dts := make([]*domain.Todo, len(todos))
	for i, todo := range todos {
		tags := todo.Edges.Tags
		if tags == nil {
			intErr := domainerrors.NewInternalServerError(("Tag is incorrect"))
			wrapErr := errors.WithStack(intErr)
			return nil, nil, wrapErr
		}

		dt, err := models.ConvertEntToTodo(todo, tags)
		if err != nil {
			return nil, nil, errors.WithStack(err)
		}
		dts[i] = dt
	}

	total, err := client.Todo.Query().Where(todo.TitleContains(title)).Count(ctx)

	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	pageMeta := domain.NewPaginationMeta(start, take, total)

	return dts, pageMeta, nil
}

func (tr *TodoRepository) FindByDescription(ctx context.Context, description string, start int, take int) ([]*domain.Todo, *domain.PaginationMeta, error) {
	if tr == nil {
		return nil, nil, errors.New("TodoRepositoryInterface pointer is nil")
	}

	client := tr.conn.GetTx(ctx)
	if client == nil {
		client = tr.conn.GetClient()
	}

	if client == nil {
		intErr := domainerrors.NewInternalServerError("client or transaction is not found")
		return nil, nil, errors.WithStack(intErr)
	}

	todos, err := client.Todo.Query().Where(todo.DescriptionContains(description)).WithTags().Offset(start).Limit(take).All(ctx)

	if err != nil && !ent.IsNotFound(err) {
		return nil, nil, errors.WithStack(err)
	}

	dts := make([]*domain.Todo, len(todos))
	for i, todo := range todos {
		tags := todo.Edges.Tags
		if tags == nil {
			intErr := domainerrors.NewInternalServerError(("Tag is incorrect"))
			wrapErr := errors.WithStack(intErr)
			return nil, nil, wrapErr
		}

		dt, err := models.ConvertEntToTodo(todo, tags)
		if err != nil {
			return nil, nil, errors.WithStack(err)
		}
		dts[i] = dt
	}

	total, err := client.Todo.Query().Where(todo.DescriptionContains(description)).Count(ctx)

	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	pageMeta := domain.NewPaginationMeta(start, take, total)

	return dts, pageMeta, nil
}

func (tr *TodoRepository) FindByImage(ctx context.Context, image string, start int, take int) ([]*domain.Todo, *domain.PaginationMeta, error) {
	if tr == nil {
		return nil, nil, errors.New("TodoRepositoryInterface pointer is nil")
	}

	client := tr.conn.GetTx(ctx)
	if client == nil {
		client = tr.conn.GetClient()
	}

	if client == nil {
		intErr := domainerrors.NewInternalServerError("client or transaction is not found")
		return nil, nil, errors.WithStack(intErr)
	}

	todos, err := client.Todo.Query().Where(todo.ImageContains(image)).WithTags().Offset(start).Limit(take).All(ctx)

	if err != nil && !ent.IsNotFound(err) {
		return nil, nil, errors.WithStack(err)
	}

	dts := make([]*domain.Todo, len(todos))
	for i, todo := range todos {
		tags := todo.Edges.Tags
		if tags == nil {
			intErr := domainerrors.NewInternalServerError(("Tag is incorrect"))
			wrapErr := errors.WithStack(intErr)
			return nil, nil, wrapErr
		}

		dt, err := models.ConvertEntToTodo(todo, tags)
		if err != nil {
			return nil, nil, errors.WithStack(err)
		}
		dts[i] = dt
	}

	total, err := client.Todo.Query().Where(todo.ImageContains(image)).Count(ctx)

	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	pageMeta := domain.NewPaginationMeta(start, take, total)

	return dts, pageMeta, nil
}

func (tr *TodoRepository) FindByTag(ctx context.Context, todotag string, start int, take int) ([]*domain.Todo, *domain.PaginationMeta, error) {
	if tr == nil {
		return nil, nil, errors.New("TodoRepositoryInterface pointer is nil")
	}

	client := tr.conn.GetTx(ctx)
	if client == nil {
		client = tr.conn.GetClient()
	}

	if client == nil {
		intErr := domainerrors.NewInternalServerError("client or transaction is not found")
		return nil, nil, errors.WithStack(intErr)
	}

	todos, err := client.Todo.Query().Where(todo.HasTagsWith(tag.Keyword(todotag))).WithTags().Offset(start).Limit(take).All(ctx)

	dts := make([]*domain.Todo, len(todos))
	for i, t := range todos {
		tags := t.Edges.Tags
		if tags == nil {
			intErr := domainerrors.NewInternalServerError(("Tag is incorrect"))
			wrapErr := errors.WithStack(intErr)
			return nil, nil, wrapErr
		}

		dt, err := models.ConvertEntToTodo(t, tags)
		if err != nil {
			return nil, nil, errors.WithStack(err)
		}
		dts[i] = dt
	}

	total, err := client.Todo.Query().Where(todo.HasTagsWith(tag.Keyword(todotag))).Count(ctx)

	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	pageMeta := domain.NewPaginationMeta(start, take, total)

	return dts, pageMeta, nil
}

func (tr *TodoRepository) FindAll(ctx context.Context, start int, take int) ([]*domain.Todo, *domain.PaginationMeta, error) {
	if tr == nil {
		return nil, nil, errors.New("TodoRepositoryInterface pointer is nil")
	}

	client := tr.conn.GetTx(ctx)
	if client == nil {
		client = tr.conn.GetClient()
	}

	if client == nil {
		intErr := domainerrors.NewInternalServerError("client or transaction is not found")
		return nil, nil, errors.WithStack(intErr)
	}

	todos, err := client.Todo.Query().WithTags().Offset(start).Limit(take).All(ctx)

	if err != nil {
		nfErr := domainerrors.NewNotFoundError("Todo", "all")
		wrapErr := errors.WithStack(nfErr)
		return nil, nil, wrapErr
	}

	dts := make([]*domain.Todo, len(todos))
	for i, todo := range todos {
		tags := todo.Edges.Tags
		if tags == nil {
			intErr := domainerrors.NewInternalServerError(("Tag is incorrect"))
			wrapErr := errors.WithStack(intErr)
			return nil, nil, wrapErr
		}

		dt, err := models.ConvertEntToTodo(todo, tags)
		if err != nil {
			return nil, nil, errors.WithStack(err)
		}

		dts[i] = dt

	}

	total, err := client.Todo.Query().Count(ctx)

	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	pageMeta := domain.NewPaginationMeta(start, take, total)

	return dts, pageMeta, nil
}

func (tr *TodoRepository) Export(ctx context.Context) ([]*domain.Todo, error) {
	if tr == nil {
		return nil, errors.New("TodoRepositoryInterface pointer is nil")
	}

	client := tr.conn.GetTx(ctx)
	if client == nil {
		client = tr.conn.GetClient()
	}

	if client == nil {
		intErr := domainerrors.NewInternalServerError("client or transaction is not found")
		return nil, errors.WithStack(intErr)
	}

	todos, err := client.Todo.Query().WithTags().All(ctx)

	if err != nil {
		nfErr := domainerrors.NewNotFoundError("Todo", "all")
		wrapErr := errors.WithStack(nfErr)
		return nil, wrapErr
	}

	dts := make([]*domain.Todo, len(todos))
	for i, todo := range todos {
		tags := todo.Edges.Tags
		if tags == nil {
			intErr := domainerrors.NewInternalServerError(("Tag is incorrect"))
			wrapErr := errors.WithStack(intErr)
			return nil, wrapErr
		}

		dt, err := models.ConvertEntToTodo(todo, tags)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		dts[i] = dt
	}

	return dts, nil
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

	client := tr.conn.GetTx(ctx)
	if client == nil {
		intErr := domainerrors.NewInternalServerError("transaction is not found")
		return nil, errors.WithStack(intErr)
	}

	todoEnt, err := client.Todo.Create().
		SetTitle(title.Value()).
		SetDescription(description.Value()).
		SetImage(image.Value()).
		SetStartsAt(startsAt).
		SetEndsAt(endsAt).
		Save(ctx)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	tagEnt, err := client.Tag.MapCreateBulk(tags, func(c *ent.TagCreate, i int) {
		c.SetKeyword(tags[i].Value()).SetTodo(todoEnt)
	}).Save(ctx)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	dt, err := models.ConvertEntToTodo(todoEnt, tagEnt)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return dt, nil
}

func (tr *TodoRepository) Delete(ctx context.Context, todoId domain.TodoId) error {
	if tr == nil {
		return errors.New("TodoRepositoryInterface pointer is nil")
	}

	client := tr.conn.GetTx(ctx)
	if client == nil {
		intErr := domainerrors.NewInternalServerError("transaction is not found")
		return errors.WithStack(intErr)
	}

	_, err := client.Tag.Delete().Where(tag.HasTodoWith(todo.ID(todoId.Value()))).Exec(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = client.Todo.Delete().Where(todo.ID(todoId.Value())).Exec(ctx)

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (tr *TodoRepository) Update(ctx context.Context,
	todoId domain.TodoId,
	title domain.TodoTitle,
	description domain.TodoDescription,
	image domain.TodoImage,
	tags []domain.TodoTag,
	startsAt time.Time,
	endsAt time.Time) (*domain.Todo, error) {

	if tr == nil {
		return nil, errors.New("TodoRepositoryInterface pointer is nil")
	}

	client := tr.conn.GetTx(ctx)
	if client == nil {
		return nil, errors.New("ent client is nil")
	}

	_, err := client.Tag.Delete().Where(tag.HasTodoWith(todo.ID(todoId.Value()))).Exec(ctx)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	todoEnt, err := client.Todo.UpdateOneID(todoId.Value()).
		SetTitle(title.Value()).
		SetDescription(description.Value()).
		SetImage(image.Value()).
		SetStartsAt(startsAt).
		SetEndsAt(endsAt).Save(ctx)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	tagsEnt, err := client.Tag.MapCreateBulk(tags, func(c *ent.TagCreate, i int) {
		c.SetKeyword(tags[i].Value()).SetTodoID(todoEnt.ID)
	}).Save(ctx)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	todo, err := models.ConvertEntToTodo(todoEnt, tagsEnt)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return todo, nil
}

func (tr *TodoRepository) CreateMany(ctx context.Context, todos []*domain.Todo) ([]*domain.Todo, error) {
	if tr == nil {
		return nil, errors.New("TodoRepositoryInterface pointer is nil")
	}

	client := tr.conn.GetTx(ctx)
	if client == nil {
		intErr := domainerrors.NewInternalServerError("transaction is not found")
		return nil, errors.WithStack(intErr)
	}

	todoCreates := make([]*ent.TodoCreate, len(todos))
	for i, todo := range todos {

		// タグの作成
		tagCreates := make([]*ent.TagCreate, len(todo.GetTags()))
		for j, tag := range todo.GetTags() {
			tagCreates[j].SetKeyword(tag.Value())
		}
		tagEnts, err := client.Tag.CreateBulk(tagCreates...).Save(ctx)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		todoCreates[i] = client.Todo.Create().
			SetTitle(todo.GetTitle().Value()).
			SetDescription(todo.GetDescription().Value()).
			SetImage(todo.GetImage().Value()).
			AddTags(tagEnts...).
			SetStartsAt(todo.GetStartsAt()).
			SetEndsAt(todo.GetEndsAt())
	}

	todoEnts, err := client.Todo.CreateBulk(todoCreates...).Save(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	dts := make([]*domain.Todo, len(todoEnts))
	for i, todoEnt := range todoEnts {
		tagEnts := todoEnt.Edges.Tags
		if tagEnts == nil {
			intErr := domainerrors.NewInternalServerError(("Tag is incorrect"))
			wrapErr := errors.WithStack(intErr)
			return nil, wrapErr
		}

		dt, err := models.ConvertEntToTodo(todoEnt, tagEnts)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		dts[i] = dt
	}
	return dts, nil
}

func (tr *TodoRepository) RunInTx(ctx context.Context, f func(context.Context) (any, error)) (any, error) {
	value, err := tr.conn.RunInTx(ctx, f)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return value, nil
}
