package src

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/tykim96/webproject/src/config"
	"github.com/tykim96/webproject/src/utils"
)

type Server struct {
	Mux *http.ServeMux

	Logger  *utils.Logger
	UTCTime *time.Time
	Config  *config.Config
	DB      *sql.DB
}

// Server start
func (s *Server) ListenAndServe(SPort uint16) error {
	err := http.ListenAndServe(fmt.Sprintf(":%d", int(SPort)), s.Mux)
	if err != nil {
		log.Fatalln("Listen Error : ", err)
		return err
	} else {
		log.Fatalln("Server Listen to SPort : ", SPort)
		return err
	}
}

// Server init
func NewServer() Server {
	return Server{
		Mux:     &http.ServeMux{},
		Logger:  utils.NewLogger(),
		UTCTime: &time.Time{},
	}
}
