package ctxhttp

import (
	"github.com/fatlotus/contextlib/reqctx"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"time"
)

// Error replies to the request with the specified error message and HTTP code.
// The error message should be plain text.
func Error(c context.Context, error string, code int) {
	w, _ := reqctx.ToRequest(c)
	http.Error(w, error, code)
}

// NotFound replies to the request with an HTTP 404 not found error.
func NotFound(c context.Context) {
	http.NotFound(reqctx.ToRequest(c))
}

// Redirect replies to the request with a redirect to url, which may
// be a path relative to the request path.
//
// The provided code should be in the 3xx range and is usually
// StatusMovedPermanently, StatusFound or StatusSeeOther.
func Redirect(c context.Context, url string, code int) {
	w, r := reqctx.ToRequest(c)
	http.Redirect(w, r, url, code)
}

// ServeFile replies to the request with the contents of the named file
// or directory.
//
// If the provided file or direcory name is a relative path, it is
// interpreted relative to the current directory and may ascend to
// parent directories. If the provided name is constructed from user
// input, it should be sanitized before calling ServeFile. As a
// precaution, ServeFile will reject requests where r.URL.Path contains
// a ".." path element.
//
// As a special case, ServeFile redirects any request where r.URL.Path
// ends in "/index.html" to the same path, without the final
// "index.html". To avoid such redirects either modify the path or use
// ServeContent.
func ServeFile(c context.Context, name string) {
	w, r := reqctx.ToRequest(c)
	http.ServeFile(w, r, name)
}

// SetCookie adds a Set-Cookie header to the provided ResponseWriter's
// headers. The provided cookie must have a valid Name. Invalid cookies
// may be silently dropped.
func SetCookie(c context.Context, cookie *http.Cookie) {
	w, _ := reqctx.ToRequest(c)
	http.SetCookie(w, cookie)
}

// ServeContent replies to the request using the content in the provided
// ReadSeeker. The main benefit of ServeContent over io.Copy is that it
// handles Range requests properly, sets the MIME type, and handles
// If-Modified-Since requests.
//
// If the response's Content-Type header is not set, ServeContent first
// tries to deduce the type from name's file extension and, if that
// fails, falls back to reading the first block of the content and
// passing it to DetectContentType. The name is otherwise unused; in
// particular it can be empty and is never sent in the response.
//
// If modtime is not the zero time or Unix epoch, ServeContent includes
// it in a Last-Modified header in the response. If the request includes
// an If-Modified-Since header, ServeContent uses modtime to decide
// whether the content needs to be sent at all.
//
// The content's Seek method must work: ServeContent uses a seek to the
// end of the content to determine its size.
//
// If the caller has set w's ETag header, ServeContent uses it to handle
// requests using If-Range and If-None-Match.
//
// Note that *os.File implements the io.ReadSeeker interface.
func ServeContent(c context.Context, name string, modtime time.Time, content io.ReadSeeker) {
	w, r := reqctx.ToRequest(c)
	http.ServeContent(w, r, name, modtime, content)
}

// Cookies parses and returns the HTTP cookies sent with the request.
func Cookies(c context.Context) []*http.Cookie {
	_, r := reqctx.ToRequest(c)
	return r.Cookies()
}

// PostFormValue returns the first value for the named component of the
// POST or PUT request body. URL query parameters are ignored.
// PostFormValue calls ParseMultipartForm and ParseForm if necessary and
// ignores any errors returned by these functions. If key is not
// present, PostFormValue returns the empty string.
func PostFormValue(c context.Context, key string) string {
	_, r := reqctx.ToRequest(c)
	return r.PostFormValue(key)
}

// ParseMultipartForm parses a request body as multipart/form-data. The
// whole request body is parsed and up to a total of maxMemory bytes of
// its file parts are stored in memory, with the remainder stored on
// disk in temporary files. ParseMultipartForm calls ParseForm if
// necessary. After one call to ParseMultipartForm, subsequent calls
// have no effect.
func ParseMultipartForm(c context.Context, maxMemory int64) error {
	_, r := reqctx.ToRequest(c)
	return r.ParseMultipartForm(maxMemory)
}

// Referer returns the referring URL, if sent in the request.
//
// Referer is misspelled as in the request itself, a mistake from the
// earliest days of HTTP. This value can also be fetched from the Header
// map as Header["Referer"]; the benefit of making it available as a
// method is that the compiler can diagnose programs that use the
// alternate (correct English) spelling req.Referrer() but cannot
// diagnose programs that use Header["Referrer"].
func Referer(c context.Context) string {
	_, r := reqctx.ToRequest(c)
	return r.Referer()
}

// UserAgent returns the client's User-Agent, if sent in the request.
func UserAgent(c context.Context) string {
	_, r := reqctx.ToRequest(c)
	return r.UserAgent()
}

// BasicAuth returns the username and password provided in the request's
// Authorization header, if the request uses HTTP Basic Authentication.
// See RFC 2617, Section 2.
func BasicAuth(c context.Context) (username, password string, ok bool) {
	_, r := reqctx.ToRequest(c)
	return r.BasicAuth()
}

// Cookie returns the named cookie provided in the request or http.ErrNoCookie
// if not found.
func Cookie(c context.Context, name string) (*http.Cookie, error) {
	_, r := reqctx.ToRequest(c)
	return r.Cookie(name)
}

// FormValue returns the first value for the named component of the
// query. POST and PUT body parameters take precedence over URL query
// string values. FormValue calls r.ParseMultipartForm and r.ParseForm if
// necessary and ignores any errors returned by these functions. If key
// is not present, FormValue returns the empty string. To access
// multiple values of the same key, call .ToRequest and inspect the
// *Request manually.
func FormValue(c context.Context, key string) string {
	_, r := reqctx.ToRequest(c)
	return r.FormValue(key)
}
