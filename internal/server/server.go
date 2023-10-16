package server

import (
	"context"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	log.Infoln("Run server")
	s.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Infoln("Shutdown server")
	return s.httpServer.Shutdown(ctx)
}
