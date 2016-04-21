// Context Functions are functions that take a context.Context as an argument.
// The ctxfunc package provides helpers that turn a ContextFunc into an
// http.Handler and back. Generally you'll need this when using existing
// non-Contextual middleware.
package ctxfunc
