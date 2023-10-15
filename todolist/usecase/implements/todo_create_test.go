package implements

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/k0825/go-gin-ent-sample/models"
	mock_interfaces "github.com/k0825/go-gin-ent-sample/repository/mock"
	"github.com/k0825/go-gin-ent-sample/usecase/interfaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTodoCreateUseCase_Handle(t *testing.T) {
	testCases := []struct {
		name        string
		title       string
		description string
		image       string
		tags        []string
		startsAt    time.Time
		endsAt      time.Time
		wantErr     bool
		prepareMock func(*mock_interfaces.MockTodoRepositoryInterface, string, string, string, []string, time.Time, time.Time) *TodoCreateInteractor
	}{
		{
			"正常系 全ての値が揃っている",
			"Todo 1",
			"This is a sample todo.",
			"https://example.com/image.png",
			[]string{"tag1", "tag2", "tag3"},
			time.Now().In(time.Local),
			time.Now().In(time.Local),
			false,
			func(m *mock_interfaces.MockTodoRepositoryInterface,
				title string,
				description string,
				image string,
				tags []string,
				startsAt time.Time,
				endsAt time.Time) *TodoCreateInteractor {
				todoId := models.NewTodoId(uuid.New())
				todoTitle, _ := models.NewTodoTitle(title)
				todoDescription, _ := models.NewTodoDescription(description)
				todoImage, _ := models.NewTodoImage(image)

				todoTags := make([]models.TodoTag, len(tags))
				for i, tag := range tags {
					todoTag, _ := models.NewTodoTag(tag)
					todoTags[i] = *todoTag
				}

				todoStartsAt := startsAt
				todoEndsAt := endsAt
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

				m.EXPECT().Create(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(expectedTodo, nil)
				tci := NewTodoCreateInteractor(m)
				return tci
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := mock_interfaces.NewMockTodoRepositoryInterface(ctrl)

			tci := tc.prepareMock(m, tc.title, tc.description, tc.image, tc.tags, tc.startsAt, tc.endsAt)

			request, err := interfaces.NewTodoCreateRequest(tc.title, tc.description, tc.image, tc.tags, tc.startsAt, tc.endsAt)
			require.NoError(t, err)

			response, err := tci.Handle(context.Background(), *request)

			if !tc.wantErr {
				assert.NoError(t, err)
				assert.Equal(t, tc.title, response.Todo.GetTitle().Value())
				assert.Equal(t, tc.description, response.Todo.GetDescription().Value())
				assert.Equal(t, tc.image, response.Todo.GetImage().Value())

				tags := response.Todo.GetTags()
				resultTags := make([]string, len(tags))
				for i, tag := range tags {
					resultTags[i] = tag.Value()
				}

				assert.Equal(t, tc.tags, resultTags)
				assert.Equal(t, tc.startsAt.Format(time.RFC3339), response.Todo.GetStartsAt().Format(time.RFC3339))
				assert.Equal(t, tc.endsAt.Format(time.RFC3339), response.Todo.GetEndsAt().Format(time.RFC3339))

			} else {
				assert.Error(t, err)
				assert.Nil(t, response)
			}
		})
	}
}
