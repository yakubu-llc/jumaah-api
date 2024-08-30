package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/joho/godotenv"
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
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env.local file")
	}

	cli := humacli.New(func(hooks humacli.Hooks, opts *Options) {
		opts.config()

		hooks.OnStart(func() {
			fmt.Printf("Starting server on port %d...\n", opts.Port)
		})
	})

	cli.Run()
}
