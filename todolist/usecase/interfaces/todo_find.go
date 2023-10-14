package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/k0825/go-gin-ent-sample/models"
)

type TodoFindRequest struct {
	Id models.TodoId
}

func NewTodoFindRequest(id uuid.UUID) *TodoFindRequest {
	todoId := models.NewTodoId(id)

	return &TodoFindRequest{Id: *todoId}
}

type TodoFindResponse struct {
	Todo models.Todo
}

func NewTodoFindResponse(todo models.Todo) *TodoFindResponse {
	return &TodoFindResponse{
		Todo: todo,
	}
}

type TodoFindUseCaseInterface interface {
	Handle(context.Context, TodoFindRequest) (*TodoFindResponse, error)
}
