package views

import (
	"github.com/gin-gonic/gin"
	"github.com/kubespace/pipeline-plugin/pkg/utils"
)

type View struct {
	Method  string
	Path    string
	Handler ViewHandler
}

func NewView(method, path string, handler ViewHandler) *View {
	return &View{
		Method:  method,
		Path:    path,
		Handler: handler,
	}
}

type ViewHandler func(*Context) *utils.Response

type Context struct {
	*gin.Context
}
