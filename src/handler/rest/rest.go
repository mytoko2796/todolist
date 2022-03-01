package restserver

import (
	"github.com/mytoko2796/sdk-go/stdlib/httpmux"
	"github.com/mytoko2796/sdk-go/stdlib/logger"
	parser "github.com/mytoko2796/sdk-go/stdlib/parser"
	"github.com/mytoko2796/todolist/src/business/usecase"
	"sync"
)

type REST interface {
	Serve()
}

type rest struct {
	logger logger.Logger
	parse   parser.Parser
	uc *usecase.Usecase
	httpmux httpmux.HttpMux
}

var once = &sync.Once{}
var e *rest

func Init(logger logger.Logger, parse parser.Parser, uc *usecase.Usecase, httpmux httpmux.HttpMux) REST{
	once.Do(func() {
		e = &rest{
			logger: logger,
			parse: parse,
			uc:   uc,
			httpmux: httpmux,
		}

		e.Serve()
	})

	return e
}

func (e *rest) Serve() {
	e.httpmux.HandleFunc(httpmux.GET, "/", e.Index)
	//enpoint for product
	e.httpmux.HandleFunc(httpmux.POST, "/v1/todolist",e.CreateToDoList)
}
