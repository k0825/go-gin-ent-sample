package interfaces

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/k0825/go-gin-ent-sample/models"
)

type TodoUpdateRequest struct {
	Id          models.TodoId
	Title       models.TodoTitle
	Description models.TodoDescription
	Image       models.TodoImage
	Tags        []models.TodoTag
	StartsAt    time.Time
	EndsAt      time.Time
}

func NewTodoUpdateRequest(
	id uuid.UUID,
	title string,
	description string,
	image string,
	tags []string,
	startsAt time.Time,
	endsAt time.Time) (*TodoUpdateRequest, error) {
	todoId := models.NewTodoId(id)
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
	return &TodoUpdateRequest{
		Id:          *todoId,
		Title:       *todoTitle,
		Description: *todoDescription,
		Image:       *todoImage,
		Tags:        todoTags,
		StartsAt:    startsAt,
		EndsAt:      endsAt,
	}, nil
}

type TodoUpdateResponse struct {
	Todo models.Todo
}

func NewTodoUpdateResponse(todo models.Todo) *TodoUpdateResponse {
	return &TodoUpdateResponse{
		Todo: todo,
	}
}

type TodoUpdateUseCaseInterface interface {
	Handle(context.Context, TodoUpdateRequest) (*TodoUpdateResponse, error)
}
