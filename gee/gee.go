package gee

import (
	"net/http"
)

// HandlerFunc 定义了gee的请求处理函数
type HandlerFunc func(*Context)

// Engine 实现了 http.Handler 接口
type Engine struct {
	router *router
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}

// New 创建一个新的 gee.Engine 实例
func New() *Engine {
	return &Engine{router: newRouter()}
}

// addRoute 添加路由规则
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
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
