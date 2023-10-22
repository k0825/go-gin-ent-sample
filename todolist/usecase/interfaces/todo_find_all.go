package interfaces

import (
	"context"

	"github.com/k0825/go-gin-ent-sample/models"
)

type TodoFindAllRequest struct {
	Start int
	Take  int
}

func NewTodoFindAllRequest(start int, take int) *TodoFindAllRequest {
	return &TodoFindAllRequest{
		Start: start,
		Take:  take,
	}
}

type TodoFindAllResponse struct {
	Todos          []*models.Todo
	PaginationMeta *models.PaginationMeta
}

func NewTodoFindAllResponse(todos []*models.Todo, pageMeta *models.PaginationMeta) *TodoFindAllResponse {
	return &TodoFindAllResponse{
		Todos:          todos,
		PaginationMeta: pageMeta,
	}
}

type TodoFindAllUseCaseInterface interface {
	Handle(context.Context, TodoFindAllRequest) (*TodoFindAllResponse, error)
}
