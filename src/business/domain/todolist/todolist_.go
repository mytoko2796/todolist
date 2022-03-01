package todolist

import (
	"context"
	"database/sql"
	x "github.com/mytoko2796/sdk-go/stdlib/error"
	apperr "github.com/mytoko2796/todolist/src/common/errors"

	"github.com/mytoko2796/todolist/src/business/entity"
)

func (t *todolistImpl) CreateToDoList(ctx context.Context, v entity.CreateTodoList) (entity.CreateTodoList, error){
	tx, err := t.sql0.Leader().BeginTx(ctx, `createToDoList`, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return v, x.WrapWithCode(err, apperr.CodeSQLTxBegin, "CreateToDoList")
	}

	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				t.logger.ErrorWithContext(ctx, err)
			}
		}
	}()

	tx, err = t.createSQLToDoList(tx, v)
	if err != nil {
		return v, x.Wrap(err, "CreateToDoList")
	}

	if err := tx.Commit(); err != nil {
		return v, x.WrapWithCode(err, apperr.CodeSQLTxCommit, "CreateToDoList")
	}

	return v, nil
}