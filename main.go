package main

import (
	"github.com/tykim96/webproject/src"
	"github.com/tykim96/webproject/src/handlers"
	"github.com/tykim96/webproject/src/models"
)

var (
	Server *src.Server
)

func init() {
	Server = src.NewServer()
	Server.ListenAndServe(3000)
}

func main() {
	defer Server.DB.Close()

	userHandler := handlers.UserHandler{DB: models.ConnectDB()}

}
