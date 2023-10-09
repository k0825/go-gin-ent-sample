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
)

func TestTodoFindUseCase_Handle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id := uuid.New()
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

	expectedTodo, err := models.NewTodo(
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

	mockRepo := mock_interfaces.NewMockTodoRepositoryInterface(ctrl)
	mockRepo.EXPECT().FindById(ctx, *todoId).Return(expectedTodo, nil)
	tfi := NewTodoFindByIdInteractor(mockRepo)

	request, err := interfaces.NewTodoFindRequest(id)
	assert.NoError(t, err)

	todo, err := tfi.Handle(ctx, *request)
	assert.NoError(t, err)

	assert.Equal(t, todo.Id, id)
}
