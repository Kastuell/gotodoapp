package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/kastuell/gotodoapp/internal/auth"
	"github.com/kastuell/gotodoapp/internal/config"
	"github.com/kastuell/gotodoapp/internal/database/postgres"
	"github.com/kastuell/gotodoapp/internal/hash"
	handler "github.com/kastuell/gotodoapp/internal/http"
	"github.com/kastuell/gotodoapp/internal/repository"
	"github.com/kastuell/gotodoapp/internal/server"
	"github.com/kastuell/gotodoapp/internal/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func Run(cfgPath string) {
	godotenv.Load()
	logrus.SetFormatter(new(logrus.JSONFormatter))
	cfg, err := config.Init(cfgPath)
	if err != nil {
		logrus.Error(err)

		return
	}

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     cfg.PostgresConfig.Host,
		Port:     cfg.PostgresConfig.Port,
		Username: cfg.PostgresConfig.Username,
		DBName:   cfg.PostgresConfig.DBName,
		SSLMode:  cfg.PostgresConfig.SSLMode,
		Password: cfg.PostgresConfig.Password,
	})

	if err != nil {
		logrus.Fatalf("failed on connecting db: %s", err.Error())
	}

	hasher := hash.NewSHA1Hasher(cfg.Auth.PasswordSalt)

	tokenManager, err := auth.NewManager(cfg.Auth.JWT.SigningKey)
	if err != nil {
		logrus.Error(err)

		return
	}

	repos := repository.NewRepository(db)
	services := service.NewService(service.NewServiceDeps{
		Repos:           repos,
		TokenManager:    tokenManager,
		Hasher:          hasher,
		AccessTokenTTL:  cfg.Auth.JWT.AccessTokenTTL,
		RefreshTokenTTL: cfg.Auth.JWT.RefreshTokenTTL,
	})
	handlers := handler.NewHandler(services, tokenManager)

	srv := server.NewServer(cfg, handlers.InitRoutes(cfg))

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logrus.Fatalf("error on running server stage: %s", err.Error())
		}
	}()

	logrus.Printf("server started on port: %s", cfg.HTTP.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("server shutting down")

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}
