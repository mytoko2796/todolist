package restserver

import "github.com/mytoko2796/todolist/src/business/entity"

type CreateToDoListRequest struct {
	Data CreateToDoListData `json:"data"`
}

type CreateToDoListData struct {
	ToDoList *entity.CreateTodoList `json:"to_do_list"`
}