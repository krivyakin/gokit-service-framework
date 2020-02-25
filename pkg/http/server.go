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
