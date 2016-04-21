package reqctx

import (
	"golang.org/x/net/context"
	"net/http"
)

type contextTag struct {
	r *http.Request
	w http.ResponseWriter
}

type writer struct {
	i http.ResponseWriter
	r *http.Request
	c context.Context
}

func (w *writer) Header() http.Header           { return w.i.Header() }
func (w *writer) Write(buf []byte) (int, error) { return w.i.Write(buf) }
func (w *writer) WriteHeader(c int)             { w.i.WriteHeader(c) }

type unexportedKey int

var reqRespKey unexportedKey = 0
var TemplateVarsKey unexportedKey = 1

// WithRequest adds a http.ResponseWriter and http.Request to a context.Context.
// Do not call WithRequest with a http.ResponseWriter returned from ToRequest;
// it will cause spooky action at a distance.
func WithRequest(c context.Context, w http.ResponseWriter, r *http.Request) context.Context {
	if writer, ok := w.(*writer); ok {
		if writer.r != r {
			panic("Do not call WithRequest on a ResponseWriter returned " +
				" from ToRequest.")
		}
		return writer.c
	}

	c = context.WithValue(c, reqRespKey, &contextTag{w: w, r: r})
	c = context.WithValue(c, TemplateVarsKey, make(map[string]interface{}))
	return c
}

// Extracts the http.ResponseWriter and *http.Request from a contextual
// request.
func ToRequest(c context.Context) (http.ResponseWriter, *http.Request) {
	u, ok := c.Value(reqRespKey).(*contextTag)
	if !ok {
		panic("context does not contain request or response.")
	}
	return &writer{u.w, u.r, c}, u.r
}
