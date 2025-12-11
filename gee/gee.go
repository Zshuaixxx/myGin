package gee

import (
	"fmt"
	"net/http"
)

// HandlerFunc 定义了gee的请求处理函数
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine 实现了 http.Handler 接口
type Engine struct {
	router map[string]HandlerFunc
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

// New 创建一个新的 gee.Engine 实例
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// addRoute 添加路由规则
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

// GET 添加 GET 请求路由规则
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST 添加 POST 请求路由规则
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run 启动 HTTP 服务器
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
