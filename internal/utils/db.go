package utils

import (
	"fmt"
	"news/internal/repository"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitDB(cfg *repository.DBConfig) (*sqlx.DB, error) {

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPassword))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
