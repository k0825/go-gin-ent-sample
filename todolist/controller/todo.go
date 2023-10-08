package controller

import (
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/k0825/go-gin-ent-sample/domainerrors"
	usecaseinterfaces "github.com/k0825/go-gin-ent-sample/usecase/interfaces"
)

type TodoControllerInterface interface {
	GetTodo(ctx *gin.Context) error
}

type TodoController struct {
	todoFindByIdUsecase usecaseinterfaces.TodoFindUseCaseInterface
}

func NewTodoController(todoFindByIdUsecase usecaseinterfaces.TodoFindUseCaseInterface) *TodoController {
	return &TodoController{
		todoFindByIdUsecase: todoFindByIdUsecase,
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
