package response

import (
	"net/http"

	"github.com/payfazz/go-handler"
	"github.com/payfazz/go-handler/defresponse"
	"github.com/payfazz/mainutil"
)

// InternalError .
func InternalError(err error) *handler.Response {
	mainutil.EprintTime(err)
	return defresponse.Status(http.StatusInternalServerError)
}

// Unauthorized .
func Unauthorized() *handler.Response {
	return defresponse.Status(http.StatusUnauthorized)
}

// LoginStatus .
func LoginStatus(href string, ok bool) *handler.Response {
	return defresponse.JSON(http.StatusOK, struct {
		Href string `json:"href"`
		Ok   bool   `json:"ok"`
	}{
		href,
		ok,
	})
}

// Counter .
func Counter(href string, counter int) *handler.Response {
	return defresponse.JSON(http.StatusOK, struct {
		Href    string `json:"href"`
		Counter int    `json:"counter"`
	}{
		href,
		counter,
	})
}
