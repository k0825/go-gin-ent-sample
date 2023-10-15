package controller

import (
	"strconv"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/k0825/go-gin-ent-sample/domainerrors"
	usecaseinterfaces "github.com/k0825/go-gin-ent-sample/usecase/interfaces"
)

type TodoControllerInterface interface {
	GetTodo(ctx *gin.Context)
	GetAllTodo(ctx *gin.Context)
	PostTodo(ctx *gin.Context)
}

type TodoController struct {
	todoFindByIdUsecase usecaseinterfaces.TodoFindUseCaseInterface
	todoFindAllUsecase  usecaseinterfaces.TodoFindAllUseCaseInterface
	todoCreateUseCase   usecaseinterfaces.TodoCreateUseCaseInterface
}

func NewTodoController(
	todoFindByIdUsecase usecaseinterfaces.TodoFindUseCaseInterface,
	todoFindAllUsecase usecaseinterfaces.TodoFindAllUseCaseInterface,
	todoCreateUsecase usecaseinterfaces.TodoCreateUseCaseInterface) *TodoController {
	return &TodoController{
		todoFindByIdUsecase: todoFindByIdUsecase,
		todoFindAllUsecase:  todoFindAllUsecase,
		todoCreateUseCase:   todoCreateUsecase,
	}
}

type jsonTime struct {
	time.Time
}

// 出力形式はRFC3339で指定
func (j jsonTime) format() string {
	return j.Format(time.RFC3339)
}

func (j jsonTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + j.format() + `"`), nil
}

type todoFindApiResponse struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Tags        []string  `json:"tags"`
	StartsAt    jsonTime  `json:"starts_at"`
	EndsAt      jsonTime  `json:"ends_at"`
	CreatedAt   jsonTime  `json:"created_at"`
	UpdatedAt   jsonTime  `json:"updated_at"`
}

type todoFindAllApiResponse struct {
	Todos  []todoFindApiResponse `json:"todos"`
	Page   int                   `json:"page"`
	Number int                   `json:"number"`
}

type todoCreateApiResponse struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Tags        []string  `json:"tags"`
	StartsAt    jsonTime  `json:"starts_at"`
	EndsAt      jsonTime  `json:"ends_at"`
	CreatedAt   jsonTime  `json:"created_at"`
	UpdatedAt   jsonTime  `json:"updated_at"`
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

func (tdc *TodoController) GetAllTodo(ctx *gin.Context) {
	spage := ctx.DefaultQuery("page", "1")
	snumber := ctx.DefaultQuery("number", "10")

	page, err := strconv.Atoi(spage)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "pageの形式が正しくありません"})
		return
	}

	number, err := strconv.Atoi(snumber)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "numberの形式が正しくありません"})
		return
	}

	request := usecaseinterfaces.NewTodoFindAllRequest(page, number)
	response, err := tdc.todoFindAllUsecase.Handle(ctx, *request)
	if err != nil {
		var nfErr *domainerrors.NotFoundError
		if errors.As(err, &nfErr) {
			ctx.JSON(404, gin.H{"message": "広告は登録されていません"})
		} else {
			ctx.JSON(500, gin.H{"message": "実行中にエラーが発生しました"})
		}
		return
	}

	resTodos := make([]todoFindApiResponse, len(response.Todos))
	for i, todo := range response.Todos {
		TodoTags := todo.GetTags()
		resTags := make([]string, len(TodoTags))
		for i, t := range TodoTags {
			resTags[i] = t.Value()
		}

		resTodos[i] = todoFindApiResponse{
			Id:          todo.GetId().Value(),
			Title:       todo.GetTitle().Value(),
			Description: todo.GetDescription().Value(),
			Image:       todo.GetImage().Value(),
			Tags:        resTags,
			StartsAt:    jsonTime{todo.GetStartsAt()},
			EndsAt:      jsonTime{todo.GetEndsAt()},
			CreatedAt:   jsonTime{todo.GetCreatedAt()},
			UpdatedAt:   jsonTime{todo.GetUpdatedAt()},
		}
	}

	resJson := todoFindAllApiResponse{
		Todos:  resTodos,
		Page:   request.Page,
		Number: request.Number,
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
	t, err := time.Parse(layout, str)
	if err != nil {
		return time.Time{}, errors.New("time parse error")
	}
	return t, err
}
