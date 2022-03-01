package restserver

import (
	"github.com/mytoko2796/todolist/src/business/entity"
)

type HTTPErrResp struct {
	Meta entity.Meta `json:"metadata"`
}

type HTTPEmptyResp struct {
	Meta entity.Meta `json:"metadata"`
}

type HTTPCreateToDoListResp struct {
	Meta entity.Meta
	Data  HTTPCreateToDoListData `json:"data"`
}

type HTTPCreateToDoListData struct {
	CreateToDoList entity.CreateTodoList `json:"create_to_do_list"`
}