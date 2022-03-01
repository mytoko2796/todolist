package todolist

import (
	"context"
	"github.com/mytoko2796/todolist/src/business/domain/todolist"
	"github.com/mytoko2796/todolist/src/business/entity"
)

type UseCaseItf interface {
	CreateToDoList(ctx context.Context, v entity.CreateTodoList) (entity.CreateTodoList, error)
}

type toDoListImpl struct {
	todoListDom todolist.DomainItf
}

func Init(todoListDom todolist.DomainItf) UseCaseItf {
	return &toDoListImpl{
		todoListDom:todoListDom,
	}
}