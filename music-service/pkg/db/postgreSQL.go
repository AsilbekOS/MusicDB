package db

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func PostgreSQLConn() (*sql.DB, error) {
	driverName := "postgres"
	dataSourceName := os.Getenv("dataSourceName")
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Println("ERROR: postgresni ochishda xatolik")
		return nil, errors.New("ERROR: postgresni ochishda xatolik")
	}

	err = db.Ping()
	if err != nil {
		log.Println("ERROR: postgresga ulanishda xatolik")
		return nil, errors.New("ERROR: postgresga ulanishda xatolik")
	}

	log.Println("PostgreSQLga muvaffaqiyatli ulandi!")
	return db, nil
}
