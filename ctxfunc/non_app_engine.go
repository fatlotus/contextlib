// +build !appengine

package ctxfunc

import (
	"golang.org/x/net/context"
	"net/http"
)

func defaultCtx(w http.ResponseWriter, r *http.Request) context.Context {
	return context.Background()
}
