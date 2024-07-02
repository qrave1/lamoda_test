package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	app2 "github.com/qrave1/lamoda_test/cmd/app"
	"github.com/qrave1/lamoda_test/cmd/commands"
	_ "github.com/qrave1/lamoda_test/cmd/commands/migrate"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "app",
		Action: func(c *cli.Context) error {
			configPath := c.String("config")
			specPath := c.String("spec")

			app := app2.NewApp()

			app.Run(configPath, specPath)

			sigCh := make(chan os.Signal, 1)
			signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
			<-sigCh

			app.Shutdown(10 * time.Second)

			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "./config/config.yaml",
				Usage:   "path to the config",
			},
			&cli.StringFlag{
				Name:    "spec",
				Aliases: []string{"s"},
				Value:   "./api/api.yaml",
				Usage:   "path to the openapi specification",
			},
		},
		Commands: commands.Commands,
	}

	if err := app.Run(os.Args); err != nil {
		panic(fmt.Errorf("error start app. %w", err))
	}
}
