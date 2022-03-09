package gee

import (
	"net/http"
)

type HandlerFune func(*Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

func (engine *Engine) addRoute(method, patter string, handler HandlerFune) {
	engine.router.addRoute(method, patter, handler)
}

func (engine *Engine) GET(pattern string, handle HandlerFune) {
	engine.addRoute("GET", pattern, handle)
}

func (engine *Engine) POST(pattern string, handler HandlerFune) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)
	engine.router.handle(c)
}
