package server

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env/v6"
	social "github.com/darmonlyone/my-social"
	"github.com/darmonlyone/my-social/postgres"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()
	if err := godotenv.Load(".env"); err != nil {
		logger.Fatal("die loading .env", zap.Error(err))
	}

	logger = logger.With(
		zap.Any("app", map[string]string{
			"name":       "social",
			"component":  "server",
			"instanceID": uuid.NewString(),
		}),
	)

	config, err := newConfigFromENV()
	if err != nil {
		logger.Fatal("die creating config from env", zap.Error(err))
	}

	postgresDSN := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		config.postgres.User,
		config.postgres.Password,
		config.postgres.Host,
		config.postgres.Port,
		config.postgres.DBName,
		config.postgres.SSL,
	)

	db, err := sql.Open("postgres", postgresDSN)
	if err != nil {
		logger.Fatal("die opening connection to postgres", zap.Error(err))
	}

	accountRepo := postgres.NewAccountRepo(db)
	postRepo := postgres.NewPostRepo(db)
	service := social.NewService(accountRepo, postRepo)

	httpHandler := social.NewHTTPServer(
		logger,
		service,
	)
	listenAddr := fmt.Sprintf(":%d", config.Port)
	ctx, cancel := context.WithCancel(context.Background())
	go social.ListenAndServe(ctx, listenAddr, httpHandler, logger)

	sig := make(chan os.Signal, 1)
	signal.Notify(
		sig,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	<-sig
	cancel()
	db.Close()
}

type config struct {
	Port int `env:"PORT" envDefault:"8000"`

	postgres postgres.Config
}

func newConfigFromENV() (*config, error) {
	var c config
	opts := env.Options{
		RequiredIfNoDef: true,
	}
	if err := env.Parse(&c, opts); err != nil {
		return nil, err
	}
	postgresConf, err := postgres.NewConfigFromENV()
	if err != nil {
		return nil, err
	}

	c.postgres = *postgresConf
	return &c, nil
}
