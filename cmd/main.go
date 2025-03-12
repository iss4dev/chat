package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	todo "to-do-list"
	"to-do-list/pkg/handler"
	"to-do-list/pkg/repository"
	"to-do-list/pkg/service"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title           TODO App API
// @version         1.0
// @description     API Server for TODO List
// @host            localhost:8000
// @BasePath        /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// Initialize configuration
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	// Initialize database connection
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("error occurred while initializing db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(todo.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			logrus.Fatalf("error occurred while running HTTP server: %s", err.Error())
		}
	}()

	logrus.Print("TodoApp started")

	// Graceful shutdown handling
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("TodoApp shutting down")

	// Shutdown the server
	if err = srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occurred on server shutdown: %s", err.Error())
	}

	// Close database connection
	if err = db.Close(); err != nil {
		logrus.Errorf("error occurred on DB connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
