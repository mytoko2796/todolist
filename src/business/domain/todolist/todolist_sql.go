package todolist

import (
	"github.com/mytoko2796/todolist/src/business/entity"
	"github.com/go-sql-driver/mysql"
	x "github.com/mytoko2796/sdk-go/stdlib/error"
	apperr "github.com/mytoko2796/todolist/src/common/errors"
	libsql "github.com/mytoko2796/sdk-go/stdlib/sql"
)

func (b *todolistImpl) createSQLToDoList(tx libsql.CommandTx, v entity.CreateTodoList) (libsql.CommandTx, error) {
	_, err := tx.Exec(`createSQLToDoList`, CreateToDoListQuery,
		v.ID,
		v.Name,
		v.StartOn,
		v.Status,
	)
	if err != nil {
		if mysqlError, ok := err.(*mysql.MySQLError); ok {
			// check duplicate constraint
			if mysqlError.Number == 1062 {
				return tx, x.WrapWithCode(err, apperr.CodeSQLUniqueConstraint, "createSQLToDoList")
			}
		}
		return tx, x.WrapWithCode(err, apperr.CodeSQLCreate, "createSQLToDoList")
	}

	return tx, nil
}

