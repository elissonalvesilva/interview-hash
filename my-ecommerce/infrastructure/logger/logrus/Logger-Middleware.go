package logrus

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

type LoggerLogrus struct {}

type (
	ResponseData struct {
		status int
		size   int
	}

	LoggingResponseWriter struct {
		http.ResponseWriter
		responseData        *ResponseData
	}
)

func (r *LoggingResponseWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b)
	r.responseData.size += size
	return size, err
}

func (r *LoggingResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
	r.responseData.status = statusCode
}

func NewLogger() *LoggerLogrus {
	return &LoggerLogrus{}
}

func (lg *LoggerLogrus) WithLogging(h http.Handler) http.Handler {
	loggingFn := func(rw http.ResponseWriter, req *http.Request) {
		start := time.Now()

		responseData := &ResponseData{
			status: 0,
			size:   0,
		}
		lrw := LoggingResponseWriter{
			ResponseWriter: rw,
			responseData:   responseData,
		}
		h.ServeHTTP(&lrw, req)
		status := lrw.responseData.status

		if status == 0 {
			status = 200
		}

		duration := time.Since(start)
		bodyBytes, _ := ioutil.ReadAll(req.Body)
		logrudFields := logrus.Fields{
			"uri":      req.RequestURI,
			"method":   req.Method,
			"remote":   req.RemoteAddr,
			"status":   status,
			"duration": duration,
			"size":     responseData.size,
			"referer":  req.Referer(),
			"user-agent": req.UserAgent(),
			"body": string(bodyBytes),
		}

		logrus.WithFields(logrudFields).Info("request completed")
	}
	return http.HandlerFunc(loggingFn)
}
