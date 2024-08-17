package main

import (
	"log/slog"
	"net/http"
	"os"
	"url-shortener/internal/config"
	handler "url-shortener/internal/handlers"
	"url-shortener/internal/lib/logger/sl"
	"url-shortener/internal/storage/postgres"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.MustLoad()

	log := sl.SetupLogger(cfg.Env)

	log.Info("starting url-shortener", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	storage, err := postgres.Open(cfg.Dsn)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	router := gin.Default()

	handler := handler.New(storage)

	router.GET("/:alias", handler.GetUrlByAlias)

	router.DELETE("/", handler.DeleteAlias)

	router.POST("/", handler.PostUrl)

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		IdleTimeout:  cfg.IdleTimeout,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}

	log.Error("server stopped")
}
