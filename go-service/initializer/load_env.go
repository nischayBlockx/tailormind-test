package initializer

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error in Loading ENV file")
	}

	log.Println("✅ .env loaded successfully")
	log.Println("🧪 DATABASE_URL =", os.Getenv("DATABASE_URL"))

}
