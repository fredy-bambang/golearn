package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/uptrace/bun"
	"golang.org/x/time/rate"
)

// dbPingLimiter limits when we actually ping the database to at most 1/sec to
// prevent a DOS since this is an unauthenticated endpoint.
var dbPingLimiter = rate.NewLimiter(rate.Every(1*time.Second), 0)

func handleHealthz(db *bun.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if db != nil {
			if dbPingLimiter.Allow() {
				if err := db.Ping(); err != nil {
					fmt.Printf("failed to ping database: %v", err)
					http.Error(w, http.StatusText(http.StatusInternalServerError),
						http.StatusInternalServerError)
					return
				}
			}
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status": "ok"}`)
	})
}
