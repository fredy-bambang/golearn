package main

import (
	"net/http"

	"github.com/fredy-bambang/golearn/internal/http/handlers"
	"github.com/payfazz/go-handler"
	"github.com/payfazz/go-handler/defresponse"
	"github.com/payfazz/go-middleware"
	"github.com/payfazz/go-router/method"
	"github.com/payfazz/go-router/path"
	"github.com/payfazz/go-router/segment"
	// ph "github.com/fredy-bambang/golearn/internal/http/handlers"
)

func main() {
	// dbName := os.Getenv("DB_NAME")
	// dbPass := os.Getenv("DB_PASS")
	// dbHost := os.Getenv("DB_HOST")
	// dbPort := os.Getenv("DB_PORT")

	// println("this is db", dbName, dbHost, dbPass, dbPort)

	// connection, err := driver.ConnectSQL(dbHost, dbPort, "root", dbPass, dbName)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(-1)
	// }

	// pHandler := handlers.NewPostHandler(connection)
	http.ListenAndServe(":1234", root())
}

func root() http.HandlerFunc {
	return path.H{
		"api": api(),
	}.C()
}

func api() http.HandlerFunc {
	return path.H{
		"coba": middleware.C(
			segment.MustEnd,
			method.Must("GET"),
			handler.Of(handlers.Coba),
		),
		// "posts": middleware.C(
		// 	segment.MustEnd,
		// 	method.Must("GET"),
		// 	handler.Of(pHandler.Fetch),
		// ),
	}.C()
}

func coba(r *http.Request) *handler.Response {
	return defresponse.Text(200, "zzz") // we can't forget this, because it'll be compile error if there is no `return`
}
