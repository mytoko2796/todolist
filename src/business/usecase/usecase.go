package usecase

import (
	"github.com/mytoko2796/todolist/src/business/domain"
	"github.com/mytoko2796/todolist/src/business/usecase/todolist"
)

type Usecase struct {
	ToDoList todolist.UseCaseItf
}

func Init(dom *domain.Domain) *Usecase{
	return &Usecase{
		ToDoList: dom.ToDoList,
	}
}
