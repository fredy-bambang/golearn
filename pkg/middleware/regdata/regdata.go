package reqdata

import (
	"net/http"

	"github.com/payfazz/go-middleware/common/kv"
	"github.com/payfazz/stdlog"
)

type ctxKeyType struct{}

var ctxKey ctxKeyType

// Data .
type Data struct {
	// App    *app.App
	InfLog stdlog.Printer
	ErrLog stdlog.Printer
}

// Inject .
func Inject(data Data) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			kv.Set(r, ctxKey, &data)
			next(w, r)
		}
	}
}

// Get .
func Get(r *http.Request) *Data {
	return kv.MustGet(r, ctxKey).(*Data)
}
