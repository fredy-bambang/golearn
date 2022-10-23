package http

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"
)

func (s *Server) Routes() *bunrouter.Router {
	router := bunrouter.New(
		bunrouter.Use(reqlog.NewMiddleware(
			reqlog.FromEnv("BUNDEBUG"),
		)),
		bunrouter.WithNotFoundHandler(NotFoundHandler),
		bunrouter.WithMethodNotAllowedHandler(MethodNotAllowedHandler),
	)

	router.GET("/", IndexHandler)
	router.POST("/405", IndexHandler) // to test methodNotAllowedHandler

	router.WithGroup("/api", func(g *bunrouter.Group) {
		g.GET("/users/:id", DebugHandler)
		g.GET("/users/current", DebugHandler)
		g.GET("/users/*path", DebugHandler)
	})

	return router
}

func IndexHandler(w http.ResponseWriter, req bunrouter.Request) error {
	return indexTemplate().Execute(w, nil)
}

func DebugHandler(w http.ResponseWriter, req bunrouter.Request) error {
	return bunrouter.JSON(w, bunrouter.H{
		"route":  req.Route(),
		"params": req.Params().Map(),
	})
}

func NotFoundHandler(w http.ResponseWriter, req bunrouter.Request) error {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(
		w,
		"<html>BunRouter can't find a route that matches <strong>%s</strong></html>",
		req.URL.Path,
	)
	return nil
}

func MethodNotAllowedHandler(w http.ResponseWriter, req bunrouter.Request) error {
	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(
		w,
		"<html>BunRouter does have a route that matches <strong>%s</strong>, "+
			"but it does not handle method <strong>%s</strong></html>",
		req.URL.Path, req.Method,
	)
	return nil
}

var indexTmpl = `
<html>
  <h1>Welcome</h1>
  <ul>
    <li><a href="/api/users/123">/api/users/123</a></li>
    <li><a href="/api/users/current">/api/users/current</a></li>
    <li><a href="/api/users/foo/bar">/api/users/foo/bar</a></li>
    <li><a href="/404">/404</a></li>
    <li><a href="/405">/405</a></li>
  </ul>
</html>
`

func indexTemplate() *template.Template {
	return template.Must(template.New("index").Parse(indexTmpl))
}
