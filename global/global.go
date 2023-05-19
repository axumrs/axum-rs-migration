package global

import (
	"log"
	"os"

	"github.com/axumrs/axum-rs-migration/db"
	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("loading .env failed:", err)
	}

	if err := db.InitPG(os.Getenv("PG_DSN")); err != nil {
		log.Fatal("connect pg failed:", err)
	}
	if err := db.InitMySQL(os.Getenv("MYSQL_DSN")); err != nil {
		log.Fatal("connect mysql failed:", err)
	}
}
