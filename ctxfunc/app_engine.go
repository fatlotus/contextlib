// +build appengine

package ctxfunc

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"net/http"
)

func defaultCtx(w http.ResponseWriter, r *http.Request) context.Context {
	return appengine.NewContext(w, r)
}
