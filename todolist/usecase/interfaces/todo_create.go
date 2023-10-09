package interfaces

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/k0825/go-gin-ent-sample/models"
)

type TodoCreateRequest struct {
	Title       models.TodoTitle
	Description models.TodoDescription
	Image       models.TodoImage
	Tags        []models.TodoTag
	StartsAt    time.Time
	EndsAt      time.Time
}

func NewTodoCreateRequest(
	title string,
	description string,
	image string,
	tags []string,
	startsAt time.Time,
	endsAt time.Time) (*TodoCreateRequest, error) {

	todoTitle, err := models.NewTodoTitle(title)
	if err != nil {
		return nil, err
	}

	todoDescription, err := models.NewTodoDescription(description)
	if err != nil {
		return nil, err
	}

	todoImage, err := models.NewTodoImage(image)
	if err != nil {
		return nil, err
	}

	todoTags := make([]models.TodoTag, len(tags))
	for i, tag := range tags {
		todoTag, err := models.NewTodoTag(tag)

		if err != nil {
			return nil, err
		}

		todoTags[i] = *todoTag
	}
	return &TodoCreateRequest{
		Title:       *todoTitle,
		Description: *todoDescription,
		Image:       *todoImage,
		Tags:        todoTags,
		StartsAt:    startsAt,
		EndsAt:      endsAt,
	}, nil
}

type TodoCreateResponse struct {
	Id          uuid.UUID
	Title       string
	Description string
	Image       string
	Tags        []string
	StartsAt    time.Time
	EndsAt      time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTodoCreateResponse(todo models.Todo) *TodoCreateResponse {
	todoTags := todo.GetTags()
	tags := make([]string, len(todoTags))
	for i, tag := range todoTags {
		tags[i] = tag.Value()
	}

	return &TodoCreateResponse{
		Id:          todo.GetId().Value(),
		Title:       todo.GetTitle().Value(),
		Description: todo.GetDescription().Value(),
		Image:       todo.GetImage().Value(),
		Tags:        tags,
		StartsAt:    todo.GetStartsAt(),
		EndsAt:      todo.GetEndsAt(),
		CreatedAt:   todo.GetCreatedAt(),
		UpdatedAt:   todo.GetUpdatedAt(),
	}
}

type TodoCreateUseCaseInterface interface {
	Handle(context.Context, TodoCreateRequest) (*TodoCreateResponse, error)
}
