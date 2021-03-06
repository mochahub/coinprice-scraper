package main

import (
	"context"
	_ "github.com/joho/godotenv/autoload"
	"github.com/mochahub/coinprice-scraper/config"
	"github.com/mochahub/coinprice-scraper/data/psql"
	"go.uber.org/fx"
	"log"
)

func main() {
	fxApp := fx.New(
		fx.Provide(config.GetSecrets),
		fx.Provide(psql.NewDatabase),
		fx.Invoke(psql.RunMigrations),
	)
	if err := fxApp.Start(context.Background()); err != nil {
		log.Printf("ERROR STARTING APP FOR MIGRATIONS: %s", err)
	}
}
