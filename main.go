package main

import (
	"github.com/Leonardo-lucas-pereira/Api_go_biblioteca/database"
	"github.com/Leonardo-lucas-pereira/Api_go_biblioteca/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()
	server.Run()

}
