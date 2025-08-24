package main

import (
	"log"
	"os"

	"github.com/UmalatDukuev/news"
	"github.com/UmalatDukuev/news/internal/handler"
	"github.com/UmalatDukuev/news/internal/repository"
	"github.com/UmalatDukuev/news/internal/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// @title Blueprint Swagger API
// @version 1.0
// @description Swagger API for News Service.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email dukuev037@mail.ru
// @BasePath /api/v1
func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Printf("warning: could not load .env file: %s", err.Error())
		}
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	// db_pass := os.Getenv("...")
	// db.Password = db_pass
	// db_pass := os.Getenv("DB_PASSWORD")
	// repository.Config.Password = db_pass
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	defer db.Close()
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(news.Server)

	port := viper.GetString("port")
	if err := srv.Run(port, handlers.InitRoutes()); err != nil {
		log.Printf("%s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
