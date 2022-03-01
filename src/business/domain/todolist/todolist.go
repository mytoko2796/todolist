package todolist

import (
	"context"
	"github.com/mytoko2796/sdk-go/stdlib/logger"
	sqlx "github.com/mytoko2796/sdk-go/stdlib/sql"

	"github.com/mytoko2796/todolist/src/business/entity"
)

type DomainItf interface {
	CreateToDoList(ctx context.Context, v entity.CreateTodoList) (entity.CreateTodoList, error)
}

type todolistImpl struct {
	logger logger.Logger
	sql0 sqlx.SQL
}

func Init(logger logger.Logger,sql0 sqlx.SQL) DomainItf{
	return &todolistImpl{
		logger: logger,
		sql0: sql0,
	}
}
