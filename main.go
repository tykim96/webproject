package main

import (
	"github.com/tykim96/webproject/src"
)

func main() {
	Server := src.NewServer()
	Server.ListenAndServe(3000)
}
