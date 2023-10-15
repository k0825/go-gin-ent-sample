package interfaces

import (
	"context"

	"github.com/k0825/go-gin-ent-sample/models"
)

type TodoFindAllRequest struct {
	Page   int
	Number int
}

func NewTodoFindAllRequest(page int, number int) *TodoFindAllRequest {
	return &TodoFindAllRequest{
		Page:   page,
		Number: number,
	}
}

type TodoFindAllResponse struct {
	Todos []*models.Todo
}

func NewTodoFindAllResponse(todos []*models.Todo) *TodoFindAllResponse {
	return &TodoFindAllResponse{
		Todos: todos,
	}
}

type TodoFindAllUseCaseInterface interface {
	Handle(context.Context, TodoFindAllRequest) (*TodoFindAllResponse, error)
}
