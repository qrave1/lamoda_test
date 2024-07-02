package main

import (
	"fmt"
	"os"

	"github.com/qrave1/lamoda_test/cmd/commands"
	_ "github.com/qrave1/lamoda_test/cmd/commands/migrate"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "app",
		Action: func(c *cli.Context) error {
			//// TODO: app here
			//_, cleanup, err := factory.InitializeService()
			//if err != nil {
			//	log.Fatal(err)
			//}
			//defer cleanup()
			//
			//sigCh := make(chan os.Signal, 1)
			//signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
			//<-sigCh
			//
			//return nil
			return nil
		},
		Commands: commands.Commands,
	}

	if err := app.Run(os.Args); err != nil {
		panic(fmt.Errorf("error start app. %w", err))
	}
}
