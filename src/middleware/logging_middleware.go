package middleware

import (
	"net/http"
	"time"

	"github.com/tykim96/webproject/src/utils"
)

func LoggingMiddleware(next http.Handler, logger *utils.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call the next handler
		next.ServeHTTP(w, r)

		// Logging
		logger.Info.Printf("Method: %s\tPath: %s\tDuration: %s",
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	})
}
