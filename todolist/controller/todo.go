package controller

import (
	"fmt"
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
	PutTodo(ctx *gin.Context)
	DeleteTodo(ctx *gin.Context)
	SearchTodo(ctx *gin.Context)
}

type TodoController struct {
	todoFindByIdUsecase usecaseinterfaces.TodoFindUseCaseInterface
	todoFindAllUsecase  usecaseinterfaces.TodoFindAllUseCaseInterface
	todoCreateUseCase   usecaseinterfaces.TodoCreateUseCaseInterface
	todoUpdateUseCase   usecaseinterfaces.TodoUpdateUseCaseInterface
	todoDeleteUseCase   usecaseinterfaces.TodoDeleteUseCaseInterface
	todoSearchUseCase   usecaseinterfaces.TodoSearchUseCaseInterface
}

func NewTodoController(
	todoFindByIdUsecase usecaseinterfaces.TodoFindUseCaseInterface,
	todoFindAllUsecase usecaseinterfaces.TodoFindAllUseCaseInterface,
	todoCreateUsecase usecaseinterfaces.TodoCreateUseCaseInterface,
	todoUpdateUsecase usecaseinterfaces.TodoUpdateUseCaseInterface,
	todoDeleteUsecase usecaseinterfaces.TodoDeleteUseCaseInterface,
	todoSearchUsecase usecaseinterfaces.TodoSearchUseCaseInterface) *TodoController {
	return &TodoController{
		todoFindByIdUsecase: todoFindByIdUsecase,
		todoFindAllUsecase:  todoFindAllUsecase,
		todoCreateUseCase:   todoCreateUsecase,
		todoDeleteUseCase:   todoDeleteUsecase,
		todoUpdateUseCase:   todoUpdateUsecase,
		todoSearchUseCase:   todoSearchUsecase,
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

type paginationMetaResponse struct {
	Start int `json:"start"`
	Take  int `json:"take"`
	Total int `json:"total"`
}

type todoFindAllApiResponse struct {
	Items  []todoFindApiResponse  `json:"items"`
	Paging paginationMetaResponse `json:"paging"`
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

type todoCreateApiRequest struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
	Tags        []string `json:"tags"`
	StartsAt    jsonTime `json:"starts_at"`
	EndsAt      jsonTime `json:"ends_at"`
}

type todoUpdateApiResponse struct {
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

type todoUpdateApiRequest struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
	Tags        []string `json:"tags"`
	StartsAt    jsonTime `json:"starts_at"`
	EndsAt      jsonTime `json:"ends_at"`
}

type todoSearchApiResponse struct {
	Items  []todoFindApiResponse  `json:"items"`
	Paging paginationMetaResponse `json:"paging"`
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
	sstart := ctx.DefaultQuery("start", "0")
	stake := ctx.DefaultQuery("take", "10")

	start, err := strconv.Atoi(sstart)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "pageの形式が正しくありません"})
		return
	}

	take, err := strconv.Atoi(stake)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "numberの形式が正しくありません"})
		return
	}

	request := usecaseinterfaces.NewTodoFindAllRequest(start, take)
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

	resPaging := paginationMetaResponse{
		Start: response.PaginationMeta.Start,
		Take:  response.PaginationMeta.Take,
		Total: response.PaginationMeta.Total,
	}

	resJson := todoFindAllApiResponse{
		Items:  resTodos,
		Paging: resPaging,
	}

	ctx.JSON(200, resJson)
}

func (tdc *TodoController) PostTodo(ctx *gin.Context) {
	var req todoCreateApiRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"message": "リクエストの形式が正しくありません"})
		return
	}

	request, err := usecaseinterfaces.NewTodoCreateRequest(req.Title, req.Description, req.Image, req.Tags, req.StartsAt.Time, req.EndsAt.Time)

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

func (tdc *TodoController) PutTodo(ctx *gin.Context) {
	id := ctx.Param("id")

	// idをuuidに変換する
	uid, err := uuid.Parse(id)

	if err != nil {
		ctx.JSON(400, gin.H{"message": "UUID変換中にエラーが発生しました。IDの形式が正しくありません"})
		return
	}

	var req todoUpdateApiRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"message": "リクエストの形式が正しくありません"})
		return
	}

	request, err := usecaseinterfaces.NewTodoUpdateRequest(uid, req.Title, req.Description, req.Image, req.Tags, req.StartsAt.Time, req.EndsAt.Time)

	if err != nil {
		ctx.JSON(500, gin.H{"message": "リクエスト生成中にエラーが発生しました"})
		return
	}

	response, err := tdc.todoUpdateUseCase.Handle(ctx, *request)
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

	resJson := todoUpdateApiResponse{
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

func (tdc *TodoController) DeleteTodo(ctx *gin.Context) {
	id := ctx.Param("id")

	// idをuuidに変換する
	uuid, err := uuid.Parse(id)

	if err != nil {
		ctx.JSON(400, gin.H{"message": "UUID変換中にエラーが発生しました。IDの形式が正しくありません"})
	}

	request := usecaseinterfaces.NewTodoDeleteRequest(uuid)
	err = tdc.todoDeleteUseCase.Handle(ctx, *request)

	if err != nil {
		var nfErr *domainerrors.NotFoundError
		if errors.As(err, &nfErr) {
			ctx.JSON(404, gin.H{"message": "指定されたIDの広告は存在しません"})
		} else {
			ctx.JSON(500, gin.H{"message": "実行中にエラーが発生しました"})
		}
		return
	}

	ctx.JSON(204, nil)
}

func (tdc *TodoController) SearchTodo(ctx *gin.Context) {
	sstart := ctx.DefaultQuery("start", "0")
	stake := ctx.DefaultQuery("take", "10")
	title := ctx.Query("title")
	description := ctx.Query("description")
	image := ctx.Query("image")
	tag := ctx.Query("tag")

	fmt.Println(title)
	fmt.Println(description)
	fmt.Println(image)
	fmt.Println(tag)

	start, err := strconv.Atoi(sstart)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "pageの形式が正しくありません"})
		return
	}

	take, err := strconv.Atoi(stake)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "numberの形式が正しくありません"})
		return
	}

	request, err := usecaseinterfaces.NewTodoSearchRequest(title, description, image, tag, start, take)
	response, err := tdc.todoSearchUseCase.Handle(ctx, *request)
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

	resPaging := paginationMetaResponse{
		Start: response.PaginationMeta.Start,
		Take:  response.PaginationMeta.Take,
		Total: response.PaginationMeta.Total,
	}

	resJson := todoSearchApiResponse{
		Items:  resTodos,
		Paging: resPaging,
	}

	ctx.JSON(200, resJson)
}
