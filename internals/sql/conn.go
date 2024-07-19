package sql

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rpstvs/csitemsapp/internals/database"
)

var DB *database.Queries

func CreateDB() {

	godotenv.Load(".env")
	dbURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Printf("Couldnt connect to database")
	}
	DB = database.New(db)
}
