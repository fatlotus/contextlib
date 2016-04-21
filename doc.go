// Writing Middleware in Go can be annoying. Though an http.Handler can call
// another, there isn't an easy way to pass values along with it.
//
// It turns out Go has a convenient, semi-builtin type for this.
// (See https://godoc.org/golang.org/x/net/context for details.)
// Contextlib is a library that standardizes the use of context.Context for
// http Handlers, and provides a basic web framework wrapping net/http.
package contextlib
