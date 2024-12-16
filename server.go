package trustwallet

import (
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string) error {

	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 512,
	}

	return s.httpServer.ListenAndServe()
}
