package main

import (
	"log"
	"server/db"
	"server/internal/user"
	"server/internal/ws"
	"server/router"
)

func main() {
	dbConn, err := db.NewDatabase()

	if err != nil {
		log.Fatalf("Could not init db connection: %s", err)
	}

	userRepo := user.NewRepository(dbConn.GetDB())
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wsHandler)
	router.Start("0.0.0.0:8080")
}
