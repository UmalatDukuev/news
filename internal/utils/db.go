package utils

import (
	"fmt"
	"news/internal/repository"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitDB(cfg *repository.DBConfig) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPassword)
	fmt.Println("Connecting to DB with DSN:", dsn)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database:", err)
		return nil, err
	}

	fmt.Println("Database connection successful.")
	return db, nil
}
