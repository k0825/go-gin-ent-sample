package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/k0825/go-gin-ent-sample/models"
)

type TodoDeleteRequest struct {
	Id models.TodoId
}

func NewTodoDeleteRequest(id uuid.UUID) *TodoDeleteRequest {
	todoId := models.NewTodoId(id)

	return &TodoDeleteRequest{Id: *todoId}
}

type TodoDeleteUseCaseInterface interface {
	Handle(context.Context, TodoDeleteRequest) error
}
