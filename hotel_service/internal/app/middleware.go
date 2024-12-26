package app

import (
	"net/http"
	"strconv"
	"time"
)

func (app *Application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.InfoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())

		next.ServeHTTP(w, r)
	})
}

type statusResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *statusResponseWriter) WriteHeader(status_code int) {
	w.statusCode = status_code
	w.ResponseWriter.WriteHeader(status_code)
}

func (app *Application) collectMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		status_recorder := statusResponseWriter{w, 200}

		next.ServeHTTP(&status_recorder, r)

		dur := time.Since(start)
		status_code := strconv.Itoa(status_recorder.statusCode)
		app.Metrics.ResponseTime.WithLabelValues(r.RequestURI, r.Method, status_code).Observe(dur.Seconds())
		app.Metrics.RequestTotal.WithLabelValues(r.RequestURI, r.Method, status_code).Inc()
	})
}
