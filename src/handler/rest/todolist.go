package restserver

import (
	x "github.com/mytoko2796/sdk-go/stdlib/error"
	apperr "github.com/mytoko2796/todolist/src/common/errors"
	"io/ioutil"
	"net/http"
)

func (e *rest) Index(w http.ResponseWriter, r *http.Request) {
	e.httpRespSuccess(w, r, http.StatusOK, nil, nil)
}

func (e *rest) CreateToDoList(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e.httpRespError(w, r, x.WrapWithCode(err, apperr.CodeHTTPBadRequest, "ioutil.ReadAll"))
		return
	}

	var requestBody CreateToDoListRequest

	if err := e.parse.JSONParser().Unmarshal(body, &requestBody);err != nil{
		e.httpRespError(w, r, x.WrapWithCode(err, apperr.CodeHTTPBadRequest, "Unmarshal"))
		return
	}

	if requestBody.Data.ToDoList == nil {
		e.httpRespError(w, r, x.NewWithCode(apperr.CodeHTTPBadRequest, "ToDoList nil"))
		return
	}

	result, err := e.uc.ToDoList.CreateToDoList(r.Context(), *requestBody.Data.ToDoList)
	if err != nil {
		e.httpRespError(w, r, x.WrapWithCode(err, apperr.CodeHTTPBadRequest, "CreateToDoList"))
	}

	e.httpRespSuccess(w, r, http.StatusCreated, result, nil)
	return
}