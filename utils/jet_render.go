package utils

import (
	"io"

	"github.com/CloudyKit/jet/v6"
	"github.com/labstack/echo/v4"
)

type JetRenderer struct {
	Views *jet.Set
}

func (r *JetRenderer) Render(w io.Writer, name string, data interface{ any }, c echo.Context) error {
	tmpl, err := r.Views.GetTemplate(name)
	if err != nil {
		return err
	}

	vars := make(jet.VarMap)

	if viewContext, isMap := data.(map[string]interface{ any }); isMap {
		for k, v := range viewContext {
			vars.Set(k, v)
		}
		vars.Set("reverse", c.Echo().Reverse)
	}

	return tmpl.Execute(w, vars, nil)
}

func NewRenderer(views *jet.Set) *JetRenderer {
	return &JetRenderer{
		Views: views,
	}
}
