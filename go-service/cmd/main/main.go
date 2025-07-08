package main

import (
	"go-service/api/handler"
	"go-service/api/routes"
	"go-service/initializer"
	"go-service/pkg/storage"
	"log"
	"os"
)

func init() {
	initializer.LoadEnv()
}

func main() {
	databaseUrl := os.Getenv("DATABASE_URL")
	if len(databaseUrl) == 0 {
		panic("Error in Database URL")
	}
	store := &storage.Store{}
	store.CreatePostgreClient(databaseUrl)

	handler.InitHandlers(store)
	r := routes.SetupRouter()
	log.Println("Starting PDF Report Service on :8081...")
	r.Run(":8081")

}
