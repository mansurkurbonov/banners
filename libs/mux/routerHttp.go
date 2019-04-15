package mux

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// _ParamsHTTP is a httprouter implementation of Params interface.
type _ParamsHTTP struct {
	params *httprouter.Params
}

// _ParamsHTTP compile validation.
var _ Params = &_ParamsHTTP{}

// ByName returns the value of the first Param which key matches the given name. If no matching Param is found, an empty string is returned.
func (this *_ParamsHTTP) ByName(name string) string {
	return this.params.ByName(name)
}

// _UserValuesHTTP is a list of user defined variables, assignable and retrievable using a key/value pair.
type _UserValuesHTTP map[string]interface{}

// _UserValuesHTTP compile validation.
var _ UserValues = &_UserValuesHTTP{}

// Save assigns the given key a value of any type and saves it in _UserValuesHTTP.
func (this *_UserValuesHTTP) Save(key string, value interface{}) {
	(*this)[key] = value
}

// Recover retrieves a value from _UserValuesHTTP by its key name.
func (this *_UserValuesHTTP) Recover(key string) interface{} {
	return (*this)[key]
}

// _ContextHTTP is a container for request specific data.
type _ContextHTTP struct {
	w          http.ResponseWriter
	r          *http.Request
	params     *_ParamsHTTP
	userValues *_UserValuesHTTP
}

// _ContextHTTP compile validation.
var _ Context = &_ContextHTTP{}

// Response returns http.ResponseWriter interface
func (this *_ContextHTTP) Response() http.ResponseWriter {
	return this.w
}

// Request returns an instance of *http.Request.
func (this *_ContextHTTP) Request() *http.Request {
	return this.r
}

// Params returns an instance of *_ParamsHTTP.
func (this *_ContextHTTP) Params() Params {
	return this.params
}

// UserValues returns an instance of *_UserValuesHTTP.
func (this *_ContextHTTP) UserValues() UserValues {
	return this.userValues
}

// ParseJsonPayload reads json data from request body into holder.
func (this *_ContextHTTP) ParseJsonPayload(holder interface{}) error {
	return json.NewDecoder(this.Request().Body).Decode(&holder)
}

// _RouterHTTP is a http.Handler which can be used to dispatch requests to different handler functions via configurable routes.
type _RouterHTTP struct {
	router *httprouter.Router
}

// _RouterHTTP compile validation.
var _ Router = &_RouterHTTP{}

// Handle registers a new request handle and optional middleware functions with the given path and method.
func (this *_RouterHTTP) Handle(method string, path string, handler Handler, mw ...Middleware) {
	for i := range mw {
		handler = mw[i](handler)
	}

	var f = func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		var ctx Context = &_ContextHTTP{
			w:          w,
			r:          r,
			params:     &_ParamsHTTP{&params},
			userValues: &_UserValuesHTTP{},
		}

		handler(ctx)
	}

	this.router.Handle(method, path, f)
}

// GET is a shortcut for router.Handle(METHOD_GET, path, handler, mw...)
func (this *_RouterHTTP) GET(path string, handler Handler, mw ...Middleware) {
	this.Handle(http.MethodGet, path, handler, mw...)
}

// POST is a shortcut for router.Handle(METHOD_POST, path, handler, mw...)
func (this *_RouterHTTP) POST(path string, handler Handler, mw ...Middleware) {
	this.Handle(http.MethodPost, path, handler, mw...)
}

// PUT is a shortcut for router.Handle(METHOD_PUT, path, handler, mw...)
func (this *_RouterHTTP) PUT(path string, handler Handler, mw ...Middleware) {
	this.Handle(http.MethodPut, path, handler, mw...)
}

// DELETE is a shortcut for router.Handle(METHOD_DELETE, path, handler, mw...)
func (this *_RouterHTTP) DELETE(path string, handler Handler, mw ...Middleware) {
	this.Handle(http.MethodDelete, path, handler, mw...)
}

// ServeFiles serves files from the given file system root. The path must end with "/*filepath", files are then served from the local path /defined/root/dir/*filepath. For example if root is "/etc" and *filepath is "passwd", the local file "/etc/passwd" would be served.
func (this *_RouterHTTP) ServeFiles(path, root string) {
	this.router.ServeFiles(path, http.Dir(root))
}

// ServeHTTP makes the router implement the http.Handler interface.
func (this *_RouterHTTP) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	this.router.ServeHTTP(w, r)
}

// NewRouterHTTPSettings returns default configuration for a new router.
func NewRouterHTTPSettings() RouterHTTPSettings {
	return RouterHTTPSettings{
		RedirectTrailingSlash:  true,
		RedirectFixedPath:      true,
		HandleMethodNotAllowed: true,
		HandleOPTIONS:          true,
		NotFound:               nil,
		MethodNotAllowed:       nil,
		PanicHandler:           nil,
	}
}

// RouterHTTPSettings is a _RouterHTTP configuration adaptor.
type RouterHTTPSettings struct {
	// Enables automatic redirection if the current route can't be matched but a
	// handler for the path with (without) the trailing slash exists.
	// For example if /foo/ is requested but a route only exists for /foo, the
	// client is redirected to /foo with http status code 301 for GET requests
	// and 307 for all other request methods.
	RedirectTrailingSlash bool

	// If enabled, the router tries to fix the current request path, if no
	// handle is registered for it.
	// First superfluous path elements like ../ or // are removed.
	// Afterwards the router does a case-insensitive lookup of the cleaned path.
	// If a handle can be found for this route, the router makes a redirection
	// to the corrected path with status code 301 for GET requests and 307 for
	// all other request methods.
	// For example /FOO and /..//Foo could be redirected to /foo.
	// RedirectTrailingSlash is independent of this option.
	RedirectFixedPath bool

	// If enabled, the router checks if another method is allowed for the
	// current route, if the current request can not be routed.
	// If this is the case, the request is answered with 'Method Not Allowed'
	// and HTTP status code 405.
	// If no other Method is allowed, the request is delegated to the NotFound
	// handler.
	HandleMethodNotAllowed bool

	// If enabled, the router automatically replies to OPTIONS requests.
	// Custom OPTIONS handlers take priority over automatic replies.
	HandleOPTIONS bool

	// Configurable http.Handler which is called when no matching route is
	// found. If it is not set, http.NotFound is used.
	NotFound http.Handler

	// Configurable http.Handler which is called when a request
	// cannot be routed and HandleMethodNotAllowed is true.
	// If it is not set, http.Error with http.StatusMethodNotAllowed is used.
	// The "Allow" header with allowed request methods is set before the handler
	// is called.
	MethodNotAllowed http.Handler

	// Function to handle panics recovered from http handlers.
	// It should be used to generate a error page and return the http error code
	// 500 (Internal Server Error).
	// The handler can be used to keep your server from crashing because of
	// unrecovered panics.
	PanicHandler func(http.ResponseWriter, *http.Request, interface{})
}

// RouterHTTPSettings compile validation.
var _ RouterSettings = RouterHTTPSettings{}

// build creates a new instance of _RouterHTTP using the given RouterHTTPSettings.
func (this RouterHTTPSettings) build() Router {
	var r = httprouter.New()

	r.RedirectTrailingSlash = this.RedirectTrailingSlash
	r.RedirectFixedPath = this.RedirectFixedPath
	r.HandleMethodNotAllowed = this.HandleMethodNotAllowed
	r.HandleOPTIONS = this.HandleOPTIONS
	r.NotFound = this.NotFound
	r.MethodNotAllowed = this.MethodNotAllowed
	r.PanicHandler = this.PanicHandler

	return &_RouterHTTP{r}
}
