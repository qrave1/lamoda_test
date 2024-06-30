package migrate

import (
	"context"
	"log"
	"time"

	"github.com/qrave1/lamoda_test/cmd/commands"
	"github.com/urfave/cli/v2"
)

func init() {
	commands.RegisterCommand(&cli.Command{
		Name: "migrate",
		Action: func(c *cli.Context) error {
			cont, cleanup, err := factory.InitializeMigrationContainer()
			if err != nil {
				log.Fatal(err)
			}
			defer cleanup()

			goose.SetBaseFS(storage.EmbedMigrations)
			err = goose.SetDialect("postgres")
			if err != nil {
				panic(err)
			}

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
			defer cancel()

			return goose.RunContext(
				ctx,
				c.Args().First(),
				cont.DB(),
				"migrations",
				c.Args().Get(1),
			)
		},
	})
}
