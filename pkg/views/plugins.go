package views

import (
	"github.com/kubespace/pipeline-plugin/pkg/plugins"
	"github.com/kubespace/pipeline-plugin/pkg/utils"
	"github.com/kubespace/pipeline-plugin/pkg/utils/code"
	"github.com/kubespace/pipeline-plugin/pkg/views/serializers"
	"net/http"
)

type PluginViews struct {
	Views    []*View
	builder  *plugins.CodeBuilder
	releaser *plugins.Releaser
}

func NewPluginViews() *PluginViews {
	pv := &PluginViews{
		builder:  plugins.NewBuilder(),
		releaser: plugins.NewReleaser(),
	}
	pv.Views = []*View{
		NewView(http.MethodPost, "/build_code_to_image", pv.buildCodeToImage),
		NewView(http.MethodPost, "/release", pv.release),
	}
	return pv
}

func (p *PluginViews) buildCodeToImage(c *Context) *utils.Response {
	var ser serializers.BuildCodeToImageSerializer

	if err := c.ShouldBind(&ser); err != nil {
		return &utils.Response{Code: code.ParamsError, Msg: err.Error()}
	}
	return p.builder.BuildCodeToImage(&ser)
}

func (p *PluginViews) release(c *Context) *utils.Response {
	var ser serializers.ReleaseSerializer

	if err := c.ShouldBind(&ser); err != nil {
		return &utils.Response{Code: code.ParamsError, Msg: err.Error()}
	}
	return p.releaser.Release(&ser)
}
