package ctxtmpl

import (
	"encoding/json"
	"github.com/fatlotus/contextlib/reqctx"
	"golang.org/x/net/context"
	"html/template"
)

func templateVars(c context.Context) map[string]interface{} {
	vars, ok := c.Value(reqctx.TemplateVarsKey).(map[string]interface{})
	if !ok || vars == nil {
		panic("TemplateKey not set or not a map")
	}
	return vars
}

// Set sets the given variable on this context. Unlike most Context operations,
// this one changes the Context instead of returning a new one.
func Set(c context.Context, key string, value interface{}) {
	vars := templateVars(c)
	vars[key] = value
}

// Renders the given named template using the context variables from this
// template.
func Render(c context.Context, t *template.Template, name string) {
	vars := templateVars(c)
	w, _ := reqctx.ToRequest(c)
	w.Header().Set("Content-type", "text/html; enctype=utf-8")
	err := t.ExecuteTemplate(w, name, vars)
	if err != nil {
		panic(err)
	}
}

// Renders the template context variables as a JSON response. Usually this aids
// in debugging, but can also be useful for Ajax calls.
func RenderJSON(c context.Context) {
	vars := templateVars(c)
	w, _ := reqctx.ToRequest(c)
	w.Header().Set("Content-type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(vars)
}
