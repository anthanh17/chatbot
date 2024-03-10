package main

import (
	"chatbot-ai/api"
	"chatbot-ai/db"
	"chatbot-ai/util"
	"context"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	ctx := context.Background()
	conn, err := db.ConnectDB(ctx, config.DBSource)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	server, err := api.NewServer(config, conn)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
