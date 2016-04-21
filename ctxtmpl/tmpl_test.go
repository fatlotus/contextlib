package ctxtmpl_test

import (
	"github.com/fatlotus/contextlib/ctxfunc"
	"github.com/fatlotus/contextlib/ctxhttp"
	"github.com/fatlotus/contextlib/ctxtmpl"

	"fmt"
	"golang.org/x/net/context"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

// Prepare a basic HTML template.
var t = template.Must(template.New("main").Parse("<h1>{{.name}}</h1>"))

// Show the user's name in either HTML or JSON.
func MainPage(c context.Context) {
	if ctxhttp.FormValue(c, "format") == "json" {
		ctxtmpl.RenderJSON(c)
	} else {
		ctxtmpl.Render(c, t, "main")
	}
}

// Adds the name onto the current context.
func AddName(inner ctxfunc.ContextFunc) ctxfunc.ContextFunc {
	return ctxfunc.ContextFunc(func(c context.Context) {
		// Set the user's name for the template.
		ctxtmpl.Set(c, "name", "Bob Jones")

		// Run the main handler.
		inner(c)
	})
}

// In this example, we show a fairly typical use of html/template along with
// context variable injection. Here the user's name is added as a template
// variable in an outer layer, which is then shown in the template.
func Example() {
	server := httptest.NewServer(ctxfunc.ToHandler(AddName(MainPage)))
	defer server.Close()

	urls := []string{"/", "/?format=json"}

	for _, url := range urls {
		resp, _ := http.Get(server.URL + url)
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("%s: %s\n", url, string(body))
	}

	// Output:
	// /: <h1>Bob Jones</h1>
	// /?format=json: {"name":"Bob Jones"}
}
