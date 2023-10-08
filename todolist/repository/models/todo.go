package models

import (
	"time"

	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
	"github.com/k0825/go-gin-ent-sample/domainerrors"
	"github.com/k0825/go-gin-ent-sample/models"
)

type TodoModel struct {
	id          uuid.UUID
	title       string
	description string
	image       string
	tags        []string
	startsAt    time.Time
	endsAt      time.Time
	createdAt   time.Time
	updatedAt   time.Time
}

func NewTodoModel(
	id uuid.UUID,
	title string,
	description string,
	image string,
	tags []string,
	startsAt time.Time,
	endsAt time.Time,
	createdAt time.Time,
	updatedAt time.Time) *TodoModel {
	return &TodoModel{
		id:          id,
		title:       title,
		description: description,
		image:       image,
		tags:        tags,
		startsAt:    startsAt,
		endsAt:      endsAt,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}
}

func ConvertToTodoModel(todo models.Todo) *TodoModel {

	todoTags := todo.GetTags()
	tags := make([]string, len(todoTags))
	for i, todoTag := range todoTags {
		tags[i] = todoTag.Value()
	}

	return &TodoModel{
		id:          todo.GetId().Value(),
		title:       todo.GetTitle().Value(),
		description: todo.GetDescription().Value(),
		image:       todo.GetImage().Value(),
		tags:        tags,
		startsAt:    todo.GetStartsAt(),
		endsAt:      todo.GetEndsAt(),
		createdAt:   todo.GetCreatedAt(),
		updatedAt:   todo.GetUpdatedAt(),
	}
}

func (todoModel TodoModel) ConvertToTodo() (*models.Todo, error) {
	id := models.NewTodoId(todoModel.id)

	title, err := models.NewTodoTitle(todoModel.title)

	if err != nil {
		iErr := domainerrors.NewInvalidValueError("TodoTitle", todoModel.title)
		wrapErr := errors.Wrap(err, iErr.Error())

		return nil, wrapErr
	}

	description, err := models.NewTodoDescription(todoModel.description)

	if err != nil {
		iErr := domainerrors.NewInvalidValueError("TodoDescription", todoModel.description)
		wrapErr := errors.Wrap(err, iErr.Error())

		return nil, wrapErr
	}

	image, err := models.NewTodoImage(todoModel.image)

	if err != nil {
		iErr := domainerrors.NewInvalidValueError("TodoImage", todoModel.image)
		wrapErr := errors.Wrap(err, iErr.Error())

		return nil, wrapErr
	}

	tags := make([]models.TodoTag, len(todoModel.tags))
	for i, tag := range todoModel.tags {
		t, err := models.NewTodoTag(tag)

		if err != nil {
			iErr := domainerrors.NewInvalidValueError("TodoTag", todoModel.tags[i])
			wrapErr := errors.Wrap(err, iErr.Error())

			return nil, wrapErr
		}

		tags[i] = *t
	}
	startsAt := todoModel.startsAt
	endsAt := todoModel.endsAt
	createdAt := todoModel.createdAt
	updatedAt := todoModel.updatedAt

	todo, err := models.NewTodo(*id, *title, *description, *image, tags, startsAt, endsAt, createdAt, updatedAt)

	if err != nil {
		iErr := domainerrors.NewInvalidValueError("Todo", id.String())
		wrapErr := errors.Wrap(err, iErr.Error())
		return nil, wrapErr
	}

	return todo, nil
}
