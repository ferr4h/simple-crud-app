package app

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"simple-crud-app/config"
	"simple-crud-app/internal/repository"
	"simple-crud-app/internal/rest"
	"simple-crud-app/internal/service"
	"simple-crud-app/pkg/database"
)

func Run(cfg *config.Config) {
	db, err := database.NewPostgresConnection(database.ConnectionInfo{
		Host:     cfg.PG.Host,
		Port:     cfg.PG.Port,
		Username: cfg.PG.Username,
		DBName:   cfg.PG.DBName,
		SSLMode:  cfg.PG.SSLMode,
		Password: cfg.PG.Password,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	gamesRepository := repository.NewGames(db)
	gamesService := service.NewGames(gamesRepository)
	handler := rest.NewHandler(gamesService)

	srv := &http.Server{
		Addr:    cfg.HTTP.Port,
		Handler: handler.InitRouter(),
	}

	initLogs()
	log.Info("Listening on port ", cfg.HTTP.Port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func initLogs() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}
