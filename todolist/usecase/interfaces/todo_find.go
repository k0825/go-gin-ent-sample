package interfaces

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/k0825/go-gin-ent-sample/models"
)

type TodoFindRequest struct {
	Id models.TodoId
}

func NewTodoFindRequest(id uuid.UUID) (*TodoFindRequest, error) {
	todoId := models.NewTodoId(id)

	return &TodoFindRequest{Id: *todoId}, nil
}

type TodoFindResponse struct {
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

func NewTodoFindResponse(todo models.Todo) *TodoFindResponse {
	todoTags := todo.GetTags()
	tags := make([]string, len(todoTags))
	for i, tag := range todoTags {
		tags[i] = tag.Value()
	}

	return &TodoFindResponse{
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

type TodoFindUseCaseInterface interface {
	Handle(context.Context, TodoFindRequest) (*TodoFindResponse, error)
}
