package main

import (
	"flag"
	"replication/database"
	"replication/server/handler"
	"replication/server/repository"
	"replication/server/server"
	"replication/utils"
)

func main() {
	port := flag.String("port", "8080", "./start -port [8000-8999] # Starts new server node at given port. Should be between 8000-8999 range")
	flag.Parse()

	err := utils.IsPortValid(*port)
	utils.Must(utils.WithPrefix("starting server", err))

	server := server.NewServer()
	db := database.CreateDB()
	repo := repository.NewReopsitory(db)
	hand := handler.NewHandler(repo)
	hand.RegisterDefaultRoutes(
		server.NewRouteGroup("/api"),
	)

	server.Start(*port)
}
