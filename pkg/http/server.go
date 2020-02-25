package http

import (
	"fmt"
	"github.com/krivyakin/gokit-service-framework/pkg/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	handler  http.Handler
	logger   log.Logger
	httpAddr string
}

type HTTPMiddleware func(next http.Handler) http.Handler

func NewServer(handler http.Handler, logger log.Logger, httpAddr string) *Server {
	logger = logger.WithLocation("http.Server")
	return &Server{
		handler:  handler,
		logger:   logger,
		httpAddr: httpAddr,
	}
}

func (s *Server) Start() error {
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		s.logger.Info("transport", "HTTP", "addr", s.httpAddr)
		server := &http.Server{
			Addr:   s. httpAddr,
			Handler: s.handler,
		}
		errs <- server.ListenAndServe()
	}()

	return <-errs
}


type LoggingResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	// WriteHeader(int) is not called if our response implicitly returns 200 OK, so
	// we default to that status code.
	return &LoggingResponseWriter{w, http.StatusOK}
}

func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.StatusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
