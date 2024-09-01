package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/yakubu-llc/jumaah-api/internal/server/http"
	"github.com/yakubu-llc/jumaah-api/internal/service/domain"
	"github.com/yakubu-llc/jumaah-api/internal/storage/postgres"

	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Options struct {
	Port        int    `help:"Port to listen on" short:"p" default:"8080"`
	DatabaseURL string `help:"Database URL" short:"d"`
	APIName     string `help:"API Name" short:"n"`
	APIVersion  string `help:"API Version" short:"v"`
	BaseURL     string `help:"Base API URL" short:"B"`
}

func (o *Options) config() {
	if port, err := strconv.Atoi(os.Getenv("PORT")); err == nil {
		o.Port = port
	}

	o.DatabaseURL = os.Getenv("DATABASE_URL")
	o.APIName = os.Getenv("API_NAME")
	o.APIVersion = os.Getenv("API_VERSION")
	o.BaseURL = os.Getenv("BASE_API_URL")
}

func main() {
	// Load environment variables from .env.local
	err := godotenv.Load(".env.local")
	if err != nil {
		fmt.Println("Error loading .env.local file")
	}

	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {
		options.config()

		ctx := context.Background()
		logger := zap.New(
			zapcore.NewCore(
				zapcore.NewJSONEncoder(zap.NewProductionConfig().EncoderConfig),
				zapcore.AddSync(os.Stdout), zap.InfoLevel))

		postgresConfig := postgres.NewConfig(options.DatabaseURL)
		repositories := postgres.NewRepository(postgresConfig, ctx, logger)

		services := domain.NewService(repositories)

		server := http.NewServer(
			services,
			options.APIName,
			options.APIVersion,
			logger,
		)

		hooks.OnStart(func() {
			fmt.Printf("Starting server on port %d...\n", options.Port)
			server.Serve(fmt.Sprintf(":%d", options.Port))
		})
	})

	cli.Run()
}
