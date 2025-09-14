package middlewares

import (
	"net/http"
)

type Middleware func(next http.Handler) http.Handler

type Manager struct {
	globalMiddleWares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddleWares: make([]Middleware, 0),
	}

}

func (mngr *Manager) Use(middlewares ...Middleware) {
	mngr.globalMiddleWares = append(mngr.globalMiddleWares, middlewares...)
}

func (mngr *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {

	h := handler

	//middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.Test)))
	for _, middleware := range middlewares {
		h = middleware(h)
	}

	// for _, middleware := range mngr.globalMiddleWares {
	// 	n = middleware(n)
	// }

	return h
}

func (mngr *Manager) WrapMux(handler http.Handler) http.Handler {

	h := handler

	//middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.Test)))
	for _, middleware := range mngr.globalMiddleWares {
		h = middleware(h)
	}
	return h
}
