package echopprof

import (
	"github.com/labstack/echo"

	"net/http"
	"net/http/pprof"
)

const textHtmlContentType = "text/html; charset=utf-8"

func Wrap(e *echo.Echo) {
	e.GET("/debug/pprof/", fromHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(echo.HeaderContentType, textHtmlContentType)
		pprof.Index(w, r)
	}).Handle)
	e.GET("/debug/pprof/heap", fromHTTPHandler(pprof.Handler("heap")).Handle)
	e.GET("/debug/pprof/goroutine", fromHTTPHandler(pprof.Handler("goroutine")).Handle)
	e.GET("/debug/pprof/block", fromHTTPHandler(pprof.Handler("block")).Handle)
	e.GET("/debug/pprof/threadcreate", fromHTTPHandler(pprof.Handler("threadcreate")).Handle)
	e.GET("/debug/pprof/cmdline", fromHandlerFunc(pprof.Cmdline).Handle)
	e.GET("/debug/pprof/profile", fromHandlerFunc(pprof.Profile).Handle)
	e.GET("/debug/pprof/symbol", fromHandlerFunc(pprof.Symbol).Handle)
	e.GET("/debug/pprof/mutex", fromHTTPHandler(pprof.Handler("mutex")).Handle)
}

var Wrapper = Wrap
