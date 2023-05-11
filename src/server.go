package src

import (
	"fmt"
	"net/http"
)

// Server struct
type Server struct {
	Mux *http.ServeMux
}

// Server init
func (s *Server) newServer() {
	s.Mux = http.NewServeMux()
}

// Server start
func (s *Server) listenAndServe(SPort uint16) error {
	return http.ListenAndServe(fmt.Sprintf(":%d", SPort), s.Mux)
}

func NewServer() Server {
	return Server{}
}
