package controller

import (
	"time"

	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/k0825/go-gin-ent-sample/domainerrors"
	usecaseinterfaces "github.com/k0825/go-gin-ent-sample/usecase/interfaces"
)

type TodoControllerInterface interface {
	GetTodo(ctx *gin.Context)
	PostTodo(ctx *gin.Context)
}

type TodoController struct {
	todoFindByIdUsecase usecaseinterfaces.TodoFindUseCaseInterface
	todoCreateUseCase   usecaseinterfaces.TodoCreateUseCaseInterface
}

func NewTodoController(todoFindByIdUsecase usecaseinterfaces.TodoFindUseCaseInterface,
	todoCreateUsecase usecaseinterfaces.TodoCreateUseCaseInterface) *TodoController {
	return &TodoController{
		todoFindByIdUsecase: todoFindByIdUsecase,
		todoCreateUseCase:   todoCreateUsecase,
	}
}

func (tdc *TodoController) GetTodo(ctx *gin.Context) {
	id := ctx.Param("id")

	// idをuuidに変換する
	uuid, err := uuid.Parse(id)

	if err != nil {
		ctx.JSON(400, gin.H{"message": "UUID変換中にエラーが発生しました。IDの形式が正しくありません"})
	}

	request := usecaseinterfaces.NewTodoFindRequest(uuid)
	response, err := tdc.todoFindByIdUsecase.Handle(ctx, *request)

	if err != nil {
		var nfErr *domainerrors.NotFoundError
		if errors.As(err, &nfErr) {
			ctx.JSON(404, gin.H{"message": "指定されたIDの広告は存在しません"})
		} else {
			ctx.JSON(500, gin.H{"message": "実行中にエラーが発生しました"})
		}
		return
	}

	TodoTags := response.Todo.GetTags()
	resTags := make([]string, len(TodoTags))
	for i, t := range TodoTags {
		resTags[i] = t.Value()
	}

	resJson := todoFindApiResponse{
		Id:          response.Todo.GetId().Value(),
		Title:       response.Todo.GetTitle().Value(),
		Description: response.Todo.GetDescription().Value(),
		Image:       response.Todo.GetImage().Value(),
		Tags:        resTags,
		StartsAt:    jsonTime{response.Todo.GetStartsAt()},
		EndsAt:      jsonTime{response.Todo.GetEndsAt()},
		CreatedAt:   jsonTime{response.Todo.GetCreatedAt()},
		UpdatedAt:   jsonTime{response.Todo.GetUpdatedAt()},
	}

	ctx.JSON(200, resJson)
}

func (tdc *TodoController) PostTodo(ctx *gin.Context) {
	title := ctx.PostForm("title")
	description := ctx.PostForm("description")
	image := ctx.PostForm("image")
	tags := ctx.PostFormArray("tags")

	startsAt, err := string2time(ctx.PostForm("starts_at"))
	if err != nil {
		ctx.JSON(400, gin.H{"message": "starts_atの形式が正しくありません"})
		return
	}

	endsAt, err := string2time(ctx.PostForm("ends_at"))
	if err != nil {
		ctx.JSON(400, gin.H{"message": "ends_atの形式が正しくありません"})
		return
	}

	request, err := usecaseinterfaces.NewTodoCreateRequest(title, description, image, tags, startsAt, endsAt)

	if err != nil {
		ctx.JSON(500, gin.H{"message": "リクエスト生成中にエラーが発生しました"})
		return
	}

	response, err := tdc.todoCreateUseCase.Handle(ctx, *request)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "実行中にエラーが発生しました"})
		return
	}

	TodoTags := response.Todo.GetTags()
	resTags := make([]string, len(TodoTags))
	for i, t := range TodoTags {
		resTags[i] = t.Value()
	}

	resJson := todoCreateApiResponse{
		Id:          response.Todo.GetId().Value(),
		Title:       response.Todo.GetTitle().Value(),
		Description: response.Todo.GetDescription().Value(),
		Image:       response.Todo.GetImage().Value(),
		Tags:        resTags,
		StartsAt:    jsonTime{response.Todo.GetStartsAt()},
		EndsAt:      jsonTime{response.Todo.GetEndsAt()},
		CreatedAt:   jsonTime{response.Todo.GetCreatedAt()},
		UpdatedAt:   jsonTime{response.Todo.GetUpdatedAt()},
	}

	ctx.JSON(200, resJson)
}

func string2time(str string) (time.Time, error) {
	layout := "2006-01-02T15:04:05Z07:00"
	return time.Parse(layout, str)
}
