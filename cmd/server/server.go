package main

import (
	"github.com/atla/dungeonsrv/pkg/server"
)

func main() {

	// TBD read in any cmd line arguments?

	server := server.NewApp()
	server.Run()
}
