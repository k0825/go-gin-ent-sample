package controller

import (
	"fmt"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/k0825/go-gin-ent-sample/domainerrors"
	usecaseinterfaces "github.com/k0825/go-gin-ent-sample/usecase/interfaces"
)

type TodoControllerInterface interface {
	GetTodo(ctx *gin.Context) error
	PostTodo(ctx *gin.Context) error
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

func (tdc *TodoController) GetTodo(ctx *gin.Context) error {
	id := ctx.Param("id")

	// idをuuidに変換する
	uuid, err := uuid.Parse(id)

	if err != nil {
		if errors.HasType(err, &domainerrors.InvalidValueError{}) {
			ctx.JSON(400, gin.H{"message": "指定されたIDの形式が正しくありません"})
		} else {
			ctx.JSON(500, gin.H{"message": "UUID変換中にエラーが発生しました"})
		}
		return err
	}

	request, err := usecaseinterfaces.NewTodoFindRequest(uuid)

	if err != nil {
		ctx.JSON(500, gin.H{"message": "リクエスト生成中にエラーが発生しました"})
		return err
	}

	response, err := tdc.todoFindByIdUsecase.Handle(ctx, *request)

	if err != nil {
		if errors.HasType(err, &domainerrors.NotFoundError{}) {
			ctx.JSON(404, gin.H{"message": "指定されたIDの広告は存在しません"})
		} else {
			ctx.JSON(500, gin.H{"message": "実行中にエラーが発生しました"})
			fmt.Println(err)
		}
		return err
	}

	resJson := todoFindApiResponse{
		Id:          response.Id,
		Title:       response.Title,
		Description: response.Description,
		Image:       response.Image,
		Tags:        response.Tags,
		StartsAt:    jsonTime{response.StartsAt},
		EndsAt:      jsonTime{response.EndsAt},
		CreatedAt:   jsonTime{response.CreatedAt},
		UpdatedAt:   jsonTime{response.UpdatedAt},
	}

	ctx.JSON(200, resJson)

	return nil
}

func (tdc *TodoController) PostTodo(ctx *gin.Context) error {
	title := ctx.PostForm("title")
	description := ctx.PostForm("description")
	image := ctx.PostForm("image")
	tags := ctx.PostFormArray("tags")

	startsAt, err := string2time(ctx.PostForm("starts_at"))
	if err != nil {
		ctx.JSON(400, gin.H{"message": "starts_atの形式が正しくありません"})
		return err
	}

	endsAt, err := string2time(ctx.PostForm("ends_at"))
	if err != nil {
		ctx.JSON(400, gin.H{"message": "ends_atの形式が正しくありません"})
		return err
	}

	request, err := usecaseinterfaces.NewTodoCreateRequest(title, description, image, tags, startsAt, endsAt)

	if err != nil {
		ctx.JSON(500, gin.H{"message": "リクエスト生成中にエラーが発生しました"})
		return err
	}

	response, err := tdc.todoCreateUseCase.Handle(ctx, *request)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "実行中にエラーが発生しました"})
		fmt.Println(err)
		return err
	}

	resJson := todoCreateApiResponse{
		Id:          response.Id,
		Title:       response.Title,
		Description: response.Description,
		Image:       response.Image,
		Tags:        response.Tags,
		StartsAt:    jsonTime{response.StartsAt},
		EndsAt:      jsonTime{response.EndsAt},
		CreatedAt:   jsonTime{response.CreatedAt},
		UpdatedAt:   jsonTime{response.UpdatedAt},
	}

	ctx.JSON(200, resJson)

	return nil
}

func string2time(str string) (time.Time, error) {
	layout := "2006-01-02T15:04:05Z07:00"
	return time.Parse(layout, str)
}
