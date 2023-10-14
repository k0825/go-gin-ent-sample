package interfaces

import (
	"context"
	"time"

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
	Todo models.Todo
}

func NewTodoCreateResponse(todo models.Todo) *TodoCreateResponse {

	return &TodoCreateResponse{
		Todo: todo,
	}
}

type TodoCreateUseCaseInterface interface {
	Handle(context.Context, TodoCreateRequest) (*TodoCreateResponse, error)
}
