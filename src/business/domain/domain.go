package domain

import (
	sqlx "github.com/mytoko2796/sdk-go/stdlib/sql"
	"github.com/mytoko2796/todolist/src/business/domain/todolist"
	log "github.com/mytoko2796/sdk-go/stdlib/logger"
)

type Domain struct {
	ToDoList todolist.DomainItf
}

func Init(
		logger log.Logger,
		sqlClient0 sqlx.SQL,
	) *Domain{
	return &Domain{
		ToDoList: todolist.Init(
				logger,
				sqlClient0,
			),
	}
}
