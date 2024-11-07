package main

import (
	"context"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ruohuaii/goddd/infrastructure/di"
	"github.com/ruohuaii/goddd/interfaces/server"
)

func main() {
	ctx := context.Background()
	container := di.NewContainer()
	err := container.Init(ctx)
	if err != nil {
		log.Fatalln("Error initializing container:", err)
	}
	defer container.Close()

	router := server.NewRouter()
	log.Fatalln(router.Run(":8766"))
}
