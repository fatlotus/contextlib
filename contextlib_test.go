package contextlib_test

import (
	"fmt"
	"github.com/fatlotus/contextlib/ctxfunc"
	"github.com/fatlotus/contextlib/reqctx"
	"golang.org/x/net/context"
	"net/http"
	"net/http/httptest"
)

// Take the "user" context variable and echo it to the screen.
func App(c context.Context) {
	w, _ := reqctx.ToRequest(c)
	fmt.Fprintf(w, "user is %s", c.Value("user"))
}

// Wrap the provided handler with a single context variable.
// A cleaner implementation would provide getters for state variables.
func Middleware(inner ctxfunc.ContextFunc) ctxfunc.ContextFunc {
	return func(c context.Context) {
		inner(context.WithValue(c, "user", "fred"))
	}
}

// A simple example of middleware passing values to an application.
func Example() {
	handler := ctxfunc.ToHandler(Middleware(App))

	w := httptest.NewRecorder()
	r := &http.Request{}
	handler.ServeHTTP(w, r)

	fmt.Printf("response: %s\n", w.Body.String())
	// Output: response: user is fred
}
