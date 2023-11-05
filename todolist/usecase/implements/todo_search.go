package implements

import (
	"context"
	"fmt"

	"github.com/cockroachdb/errors"
	repositoryinterfaces "github.com/k0825/go-gin-ent-sample/repository/interfaces"
	"github.com/k0825/go-gin-ent-sample/usecase/interfaces"
)

type TodoSearchInteractor struct {
	TodoRepository repositoryinterfaces.TodoRepositoryInterface
}

func NewTodoSearchInteractor(tr repositoryinterfaces.TodoRepositoryInterface) *TodoSearchInteractor {
	return &TodoSearchInteractor{
		TodoRepository: tr,
	}
}

func (tsi *TodoSearchInteractor) Handle(ctx context.Context, request interfaces.TodoSearchRequest) (*interfaces.TodoSearchResponse, error) {
	if tsi == nil {
		return nil, errors.New("TodoSearchInteractor is nil.")
	}

	response := &interfaces.TodoSearchResponse{}

	if request.Title != "" {
		fmt.Println("titleの検索")
		todos, pageMeta, err := tsi.TodoRepository.FindByTitle(ctx, request.Title, request.Start, request.Take)
		if err != nil {
			return nil, err
		}
		response = interfaces.NewTodoSearchResponse(todos, pageMeta)

	} else if request.Description != "" {
		fmt.Println("descriptionの検索")
		todos, pageMeta, err := tsi.TodoRepository.FindByDescription(ctx, request.Description, request.Start, request.Take)
		if err != nil {
			return nil, err
		}
		response = interfaces.NewTodoSearchResponse(todos, pageMeta)

	} else if request.Image != "" {
		fmt.Println("imageの検索")
		todos, pageMeta, err := tsi.TodoRepository.FindByImage(ctx, request.Image, request.Start, request.Take)
		if err != nil {
			return nil, err
		}
		response = interfaces.NewTodoSearchResponse(todos, pageMeta)

	} else if request.Tag != "" {
		fmt.Println("tagの検索")
		todos, pageMeta, err := tsi.TodoRepository.FindByTag(ctx, request.Tag, request.Start, request.Take)
		if err != nil {
			return nil, err
		}
		response = interfaces.NewTodoSearchResponse(todos, pageMeta)

	} else {
		todos, pageMeta, err := tsi.TodoRepository.FindAll(ctx, request.Start, request.Take)
		if err != nil {
			return nil, err
		}
		response = interfaces.NewTodoSearchResponse(todos, pageMeta)
	}
	return response, nil
}
