package todolist

import (
	"context"

	"github.com/mytoko2796/todolist/src/business/entity"
)

func (t *toDoListImpl) CreateToDoList(ctx context.Context, v entity.CreateTodoList) (entity.CreateTodoList, error){
	return t.todoListDom.CreateToDoList(ctx, v)
}
