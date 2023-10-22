package implements

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/k0825/go-gin-ent-sample/ent"
	"github.com/k0825/go-gin-ent-sample/ent/enttest"
	"github.com/k0825/go-gin-ent-sample/ent/migrate"
	"github.com/k0825/go-gin-ent-sample/mock_datasource"
	domain "github.com/k0825/go-gin-ent-sample/models"
	"github.com/k0825/go-gin-ent-sample/repository/models"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTodoRepository_GetTodo_Normal(t *testing.T) {
	testCases := []struct {
		name          string
		setDummyFn    func(client *ent.Client) *domain.Todo
		prepareMockFn func(m *mock_datasource.MockRDBConnectionInterface, client *ent.Client) (*TodoRepository, error)
	}{
		{
			name: "正常系 全ての値が揃っている",
			setDummyFn: func(client *ent.Client) *domain.Todo {
				client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true))

				todo, _ := client.Todo.Create().
					SetTitle("Todo 1").
					SetDescription("This is a sample todo.").
					SetImage("https://example.com/image.png").
					SetStartsAt(time.Now().In(time.Local)).
					SetEndsAt(time.Now().In(time.Local)).
					Save(context.Background())

				tagKeywords := []string{"sample", "test", "todo"}

				tags, _ := client.Tag.MapCreateBulk(tagKeywords, func(c *ent.TagCreate, i int) {
					c.SetKeyword(tagKeywords[i]).
						SetTodoID(todo.ID)
				}).Save(context.Background())

				dt, _ := models.ConvertEntToTodo(todo, tags)

				return dt
			},
			prepareMockFn: func(m *mock_datasource.MockRDBConnectionInterface, client *ent.Client) (*TodoRepository, error) {
				m.EXPECT().GetTx(context.Background()).Return(client)
				// m.EXPECT().GetClient().Return(client)
				repo, err := NewTodoRepository(m)

				if err != nil {
					return nil, err
				}
				return repo, nil
			},
		},
		{
			name: "正常系 画像がない",
			setDummyFn: func(client *ent.Client) *domain.Todo {
				client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true))

				todo, _ := client.Todo.Create().
					SetTitle("Todo 1").
					SetDescription("This is a sample todo.").
					SetImage("").
					SetStartsAt(time.Now().In(time.Local)).
					SetEndsAt(time.Now().In(time.Local)).
					Save(context.Background())

				tagKeywords := []string{"sample", "test", "todo"}

				tags, _ := client.Tag.MapCreateBulk(tagKeywords, func(c *ent.TagCreate, i int) {
					c.SetKeyword(tagKeywords[i]).
						SetTodoID(todo.ID)
				}).Save(context.Background())

				dt, _ := models.ConvertEntToTodo(todo, tags)

				return dt
			},
			prepareMockFn: func(m *mock_datasource.MockRDBConnectionInterface, client *ent.Client) (*TodoRepository, error) {
				m.EXPECT().GetTx(context.Background()).Return(client)
				repo, err := NewTodoRepository(m)

				if err != nil {
					return nil, err
				}
				return repo, nil
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := mock_datasource.NewMockRDBConnectionInterface(ctrl)

			client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			defer client.Close()

			expected := tc.setDummyFn(client)
			repo, err := tc.prepareMockFn(m, client)
			require.NoError(t, err)

			todoId := domain.NewTodoId(expected.GetId().Value())
			result, err := repo.FindById(context.Background(), *todoId)
			assert.NoError(t, err)

			assert.Equal(t, expected.GetTitle(), result.GetTitle())
			assert.Equal(t, expected.GetDescription(), result.GetDescription())
			assert.Equal(t, expected.GetImage(), result.GetImage())
			assert.Equal(t, expected.GetTags(), result.GetTags())
			assert.Equal(t, expected.GetStartsAt().Format(time.RFC3339), result.GetStartsAt().Format(time.RFC3339))
			assert.Equal(t, expected.GetEndsAt().Format(time.RFC3339), result.GetEndsAt().Format(time.RFC3339))
		})
	}
}

func TestTodoRepository_GetTodo_Abnormal(t *testing.T) {
	testCases := []struct {
		name          string
		todoId        string
		prepareMockFn func(m *mock_datasource.MockRDBConnectionInterface, client *ent.Client) (*TodoRepository, error)
	}{
		{
			name:   "異常系 存在しないIDを指定した",
			todoId: "a0cff1cc-475a-4a18-9997-2c6363f96236",
			prepareMockFn: func(m *mock_datasource.MockRDBConnectionInterface, client *ent.Client) (*TodoRepository, error) {
				m.EXPECT().GetTx(context.Background()).Return(client)
				repo, err := NewTodoRepository(m)

				if err != nil {
					return nil, err
				}
				return repo, nil
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			defer client.Close()
			client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true))

			m := mock_datasource.NewMockRDBConnectionInterface(ctrl)
			repo, err := tc.prepareMockFn(m, client)
			require.NoError(t, err)

			todoUuid := uuid.MustParse(tc.todoId)
			todoId := domain.NewTodoId(todoUuid)
			_, err = repo.FindById(context.Background(), *todoId)
			require.Error(t, err)
		})
	}
}
