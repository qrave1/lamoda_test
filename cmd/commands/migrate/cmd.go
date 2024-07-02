package migrate

import (
	"context"
	"time"

	"github.com/pressly/goose/v3"
	"github.com/qrave1/lamoda_test/cmd/commands"
	"github.com/qrave1/lamoda_test/config"
	"github.com/qrave1/lamoda_test/internal/infrastructure/persistence/postgres"
	"github.com/qrave1/lamoda_test/migrations"
	"github.com/qrave1/lamoda_test/pkg/logger"
	"github.com/urfave/cli/v2"
)

func init() {
	commands.RegisterCommand(&cli.Command{
		Name: "migrate",
		Action: func(c *cli.Context) error {
			log := logger.New()

			configPath := c.String("config")

			cfg, err := config.ReadConfig(configPath)
			if err != nil {
				log.Error("error read config", "error", err)
			}

			conn, err := postgres.NewConnect(cfg)
			if err != nil {
				log.Error("error connect postgres", "error", err)
			}
			defer conn.Close()

			goose.SetBaseFS(migrations.EmbedMigrations)
			err = goose.SetDialect("postgres")
			if err != nil {
				panic(err)
			}

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
			defer cancel()

			return goose.RunContext(
				ctx,
				c.Args().First(),
				conn,
				".",
				c.Args().Get(1),
			)
		},
	})
}
