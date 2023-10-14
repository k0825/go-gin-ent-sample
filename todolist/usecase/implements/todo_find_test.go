package implements

import (
	"context"
	"testing"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/k0825/go-gin-ent-sample/models"
	mock_interfaces "github.com/k0825/go-gin-ent-sample/repository/mock"
	"github.com/k0825/go-gin-ent-sample/usecase/interfaces"
	"github.com/stretchr/testify/assert"
)

func TestTodoFindUseCase_Handle(t *testing.T) {
	testCases := []struct {
		name        string
		wantErr     bool
		prepareMock func(*mock_interfaces.MockTodoRepositoryInterface, uuid.UUID) *TodoFindByIdInteractor
	}{
		{
			"正常系 全ての値が揃っている",
			false,
			func(m *mock_interfaces.MockTodoRepositoryInterface, id uuid.UUID) *TodoFindByIdInteractor {
				todoId := models.NewTodoId(id)
				todoTitle, _ := models.NewTodoTitle("Todo 1")
				todoDescription, _ := models.NewTodoDescription("This is a sample todo.")
				todoImage, _ := models.NewTodoImage("https://example.com/image.png")

				tags := []string{"tag1", "tag2", "tag3"}
				todoTags := make([]models.TodoTag, len(tags))
				for i, tag := range tags {
					todoTag, _ := models.NewTodoTag(tag)
					todoTags[i] = *todoTag
				}

				todoStartsAt := time.Now().In(time.Local)
				todoEndsAt := time.Now().In(time.Local)
				todoCreatedAt := time.Now().In(time.Local)
				todoUpdatedAt := time.Now().In(time.Local)

				expectedTodo, _ := models.NewTodo(
					*todoId,
					*todoTitle,
					*todoDescription,
					*todoImage,
					todoTags,
					todoStartsAt,
					todoEndsAt,
					todoCreatedAt,
					todoUpdatedAt,
				)

				ctx := context.Background()

				m.EXPECT().FindById(ctx, *todoId).Return(expectedTodo, nil)
				tfi := NewTodoFindByIdInteractor(m)
				return tfi
			},
		},
		{
			"異常系 なんらかのエラーが発生した場合、エラーを返す",
			true,
			func(m *mock_interfaces.MockTodoRepositoryInterface, id uuid.UUID) *TodoFindByIdInteractor {
				todoId := models.NewTodoId(id)

				ctx := context.Background()

				err := errors.New("some error")

				m.EXPECT().FindById(ctx, *todoId).Return(nil, err)
				tfi := NewTodoFindByIdInteractor(m)
				return tfi
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			id := uuid.New()

			m := mock_interfaces.NewMockTodoRepositoryInterface(ctrl)

			tfi := tc.prepareMock(m, id)

			request := interfaces.NewTodoFindRequest(id)

			response, err := tfi.Handle(context.Background(), *request)

			if !tc.wantErr {
				assert.NoError(t, err)
				assert.Equal(t, response.Todo.GetId().Value(), id)
			} else {
				assert.Error(t, err)
				assert.Nil(t, response)
			}
		})
	}
}
