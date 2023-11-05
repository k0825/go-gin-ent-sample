package interfaces

import (
	"context"

	"github.com/k0825/go-gin-ent-sample/models"
)

type TodoSearchRequest struct {
	Title       string
	Description string
	Image       string
	Tag         string
	Start       int
	Take        int
}

func NewTodoSearchRequest(title string, description string, image string, tag string, start int, take int) (*TodoSearchRequest, error) {

	return &TodoSearchRequest{Title: title, Description: description, Image: image, Tag: tag, Start: start, Take: take}, nil
}

type TodoSearchResponse struct {
	Todos          []*models.Todo
	PaginationMeta *models.PaginationMeta
}

func NewTodoSearchResponse(todos []*models.Todo, pageMeta *models.PaginationMeta) *TodoSearchResponse {
	return &TodoSearchResponse{
		Todos:          todos,
		PaginationMeta: pageMeta,
	}
}

type TodoSearchUseCaseInterface interface {
	Handle(context.Context, TodoSearchRequest) (*TodoSearchResponse, error)
}
