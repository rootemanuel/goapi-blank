package service

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"

	. "github.com/rootemanuel/goapi-blank/dao"
)

type UserService struct{}

func (m *UserService) Teste(ctx context.Context) {
	dao := RootDao{}

	rtn, err := dao.FindAll()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(rtn)
}
