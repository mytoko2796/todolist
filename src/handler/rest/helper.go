package restserver

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	apperr "github.com/mytoko2796/todolist/src/common/errors"
	"github.com/mytoko2796/todolist/src/common/stdlib"
	"github.com/mytoko2796/todolist/src/business/entity"
	"time"
)

func (e *rest) httpRespError(w http.ResponseWriter, r *http.Request, err error) {
	debugMode := false

	if r.Header.Get(httpheader.AppDebug) == "true"{
		debugMode = true
	}

	statusCode, displayError := apperr.CompileError(err,"", debugMode)
	statusStr := http.StatusText(statusCode)

	if statusCode == http.StatusInternalServerError {
		errCtx := r.Context().Err()
		if errCtx != nil {
			if strings.Contains(errCtx.Error(), "context canceled") {
				// 499 Client Closed Request
				statusCode = 499
				statusStr = "Client Closed Request"
			}
		}
	}

	jsonErrResp := &HTTPErrResp{
		Meta: entity.Meta{
			Path:       r.URL.String(),
			StatusCode: statusCode,
			Status:     statusStr,
			Message:    fmt.Sprintf("%s %s [%d] %s", r.Method, r.URL.RequestURI(), statusCode, statusStr),
			Error:      displayError,
			Timestamp:  time.Now().Format(time.RFC3339),
		},
	}

	raw, err := e.parse.JSONParser().Marshal(jsonErrResp)
	if err != nil {
		statusCode = http.StatusInternalServerError
	}

	w.Header().Set(httpheader.ContentType, httpheader.ContentJSON)
	w.WriteHeader(statusCode)
	_, _ = w.Write(raw)
}



func (e *rest) httpRespSuccess(w http.ResponseWriter, r *http.Request, statusCode int, resp interface{}, p *entity.Pagination) {

	meta := entity.Meta{
		Path:       r.URL.String(),
		StatusCode: statusCode,
		Status:     http.StatusText(statusCode),
		Message:    fmt.Sprintf("%s %s [%d] %s", r.Method, r.URL.RequestURI(), statusCode, http.StatusText(statusCode)),
		Error:      nil,
		Timestamp:  time.Now().Format(time.RFC3339),
	}

	var (
		schemaName string
		raw        []byte
		err        error
	)

	switch data := resp.(type) {
	case nil:
		httpResp := &HTTPEmptyResp{
			Meta: meta,
		}
		raw, err = e.Marshal("", httpResp)

	case entity.CreateTodoList:
		schemaName = ""
		checkinManifestResp := &HTTPCreateToDoListResp{
			Meta: meta,
			Data: HTTPCreateToDoListData{
				CreateToDoList: data,
			},
		}
		raw, err = e.Marshal(schemaName, checkinManifestResp)

	default:
		e.httpRespError(w, r, errors.New(fmt.Sprintf("cannot cast type of %+v", data)))
		return
	}

	if err != nil {
		e.httpRespError(w, r, errors.New("MarshalHTTPResp"))
		return
	}

	w.Header().Set(httpheader.ContentType, httpheader.ContentJSON)
	w.WriteHeader(statusCode)
	_, _ = w.Write(raw)
}

func (e *rest) Marshal(schema string, resp interface{}) ([]byte, error) {
	return e.parse.JSONParser().Marshal(&resp)
}

