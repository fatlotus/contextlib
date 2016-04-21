package ctxfunc

import (
	"github.com/fatlotus/contextlib/reqctx"
	"golang.org/x/net/context"
	"net/http"
)

// A ContextFunc is the context.Context-enabled version of a http.HandleFunc.
//
// See the reqctx package for ways of attaching and detaching the request
// from a context.Context.
type ContextFunc func(ctx context.Context)

// ToHandler converts a Context function into a handler function. Think of this
// as converting
//
//   func A(c context.Context) {
//       // do stuff with c
//   }
//
// to
//
//   func A(w http.ResponseWriter, r *http.Request) {
//       c := reqctx.FromRequest(w, r)
//       // do stuff with c
//   }
func ToHandler(cf ContextFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cf(reqctx.WithRequest(defaultCtx(w, r), w, r))
	})
}

// Converts an http.Handler to a ctx.ContextFunc.
// FromHandler converts a handler function into a Context function. Think of
// this as converting
//
//   func A(w http.ResponseWriter, r *http.Request) {
//       // do stuff with w and r
//   }
//
// to
//
//   func A(c context.Context) {
//       w, r := reqctx.ToRequest(c)
//       // do stuff with w and r
//   }
func FromHandler(h http.Handler) ContextFunc {
	return (func(c context.Context) {
		h.ServeHTTP(reqctx.ToRequest(c))
	})
}
