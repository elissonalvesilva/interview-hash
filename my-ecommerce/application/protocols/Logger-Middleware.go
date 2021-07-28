package protocols

import "net/http"

type LoggerMiddleware interface {
	WithLogging(h http.Handler) http.Handler
}
