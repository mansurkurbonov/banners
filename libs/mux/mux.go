// mux package provides custom router functionality.
package mux

import (
	"net/http"

	"errors"
)

// Errors
const (
	errorRouterSettingsNil = "router settings cannot be nil"
)

// Params is an array of URL parameters, consisting of a key and a value.
type Params interface {
	// ByName returns the value of the first Param which key matches the given name.
	// If no matching Param is found, an empty string is returned.
	ByName(name string) string
}

// UserValues is a list of user defined variables, assignable and retrievable using a key/value pair.
type UserValues interface {
	// Save assigns the given key a value of any type and saves it in UserValues.
	Save(key string, value interface{})

	// Recover retrieves a value from UserValues by its key name.
	Recover(key string) interface{}
}

// Context is a container for request specific data.
type Context interface {
	// Response returns http.ResponseWriter interface.
	Response() http.ResponseWriter

	// Request returns an instance of *http.Request.
	Request() *http.Request

	// Params returns Params interface.
	Params() Params

	// UserValues returns UserValues interface.
	UserValues() UserValues

	// ParseJsonPayload reads json data from request body into holder.
	ParseJsonPayload(holder interface{}) error
}

// Handler is a function that can be registered to a route to handle HTTP requests.
type Handler func(Context)

// Middleware is a function that can be used to chain Handlers.
type Middleware func(Handler) Handler

// Router is a http.Handler which responds to an HTTP request.
type Router interface {
	// GET is a shortcut for router.Handle(METHOD_GET, path, handler, mw...)
	GET(path string, handler Handler, mw ...Middleware)

	// POST is a shortcut for router.Handle(METHOD_POST, path, handler, mw...)
	POST(path string, handler Handler, mw ...Middleware)

	// PUT is a shortcut for router.Handle(METHOD_PUT, path, handler, mw...)
	PUT(path string, handler Handler, mw ...Middleware)

	// DELETE is a shortcut for router.Handle(METHOD_DELETE, path, handler, mw...)
	DELETE(path string, handler Handler, mw ...Middleware)

	// Handle registers a new request handle and optional middleware functions with the given path and method.
	Handle(method string, path string, handler Handler, mw ...Middleware)

	// ServeFiles serves files from the given file system root.
	// The path must end with "/*filepath", files are then served from the local path /defined/root/dir/*filepath.
	// For example if root is "/etc" and *filepath is "passwd", the local file "/etc/passwd" would be served.
	ServeFiles(path, rootPath string)

	// ServeHTTP makes the router implement the http.Handler interface.
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// RouterSettings is a Router configuration adaptor.
type RouterSettings interface {
	// build creates a new instance of Router using the given RouterSettings.
	build() Router
}

// NewRouter creates a new instance of Router using the given configuration adaptor.
func NewRouter(settings RouterSettings) (Router, error) {
	if settings == nil {
		return nil, errors.New(errorRouterSettingsNil)
	}
	return settings.build(), nil
}
