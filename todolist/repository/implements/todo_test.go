package implements

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/k0825/go-gin-ent-sample/ent"
	"github.com/k0825/go-gin-ent-sample/ent/enttest"
	"github.com/k0825/go-gin-ent-sample/ent/migrate"
	domain "github.com/k0825/go-gin-ent-sample/models"
	"github.com/k0825/go-gin-ent-sample/repository/models"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func convertTestTodo(todo *ent.Todo, tags []*ent.Tag) (*domain.Todo, error) {
	// 長さ指定して初期化
	tagKeywords := make([]string, len(tags))

	for i, tag := range tags {
		tagKeywords[i] = tag.Keyword
	}
	mtm := models.NewTodoModel(
		todo.ID,
		todo.Title,
		todo.Description,
		*todo.Image,
		tagKeywords,
		todo.StartsAt,
		todo.EndsAt,
		todo.CreatedAt,
		todo.UpdatedAt,
	)

	dt, err := mtm.ConvertToTodo()

	return dt, err

}

func TestTodoRepository_GetTodo_Normal(t *testing.T) {
	tests := []struct {
		name        string
		title       string
		description string
		image       string
		startsAt    time.Time
		endsAt      time.Time
		tagKeywords []string
	}{
		{
			"正常系 全ての値が揃っている",
			"Todo 1",
			"This is a sample todo.",
			"https://example.com/image.png",
			time.Now().In(time.Local),
			time.Now().In(time.Local),
			[]string{"sample", "test", "todo"},
		},
		{
			"正常系 画像がない",
			"Todo 2",
			"This is a sample todo.",
			"",
			time.Now().In(time.Local),
			time.Now().In(time.Local),
			[]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			defer client.Close()
			client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true))

			// Insert sample data into the database
			todo, err := client.Todo.Create().
				SetTitle(tt.name).
				SetDescription(tt.description).
				SetImage(tt.image).
				SetStartsAt(tt.startsAt).
				SetEndsAt(tt.endsAt).Save(context.Background())
			require.NoError(t, err)

			tags, err := client.Tag.MapCreateBulk(tt.tagKeywords, func(c *ent.TagCreate, i int) {
				c.SetKeyword(tt.tagKeywords[i]).
					SetTodoID(todo.ID)
			}).Save(context.Background())

			require.NoError(t, err)

			repo, err := NewTodoRepository(client)
			require.NoError(t, err)

			todoId := domain.NewTodoId(todo.ID)
			result, err := repo.FindById(context.Background(), *todoId)
			require.NoError(t, err)

			expected, err := convertTestTodo(todo, tags)
			assert.NoError(t, err)
			assert.Equal(t, expected.GetTitle(), result.GetTitle())
			assert.Equal(t, expected.GetDescription(), result.GetDescription())
			assert.Equal(t, expected.GetImage(), result.GetImage())
			assert.Equal(t, expected.GetTags(), result.GetTags())
			reflect.DeepEqual(expected.GetStartsAt(), result.GetStartsAt())
			reflect.DeepEqual(expected.GetEndsAt(), result.GetEndsAt())
		})
	}
}

func TestTodoRepository_GetTodo_NotFound(t *testing.T) {
	tests := []struct {
		name   string
		todoId string
	}{
		{
			"異常系 存在しないIDを指定した",
			"a0cff1cc-475a-4a18-9997-2c6363f96236",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			defer client.Close()
			client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true))

			repo, err := NewTodoRepository(client)
			require.NoError(t, err)

			todoUuid := uuid.MustParse(tt.todoId)

			todoId := domain.NewTodoId(todoUuid)
			_, err = repo.FindById(context.Background(), *todoId)
			require.Error(t, err)
		})
	}
}
