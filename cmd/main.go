package main

import (
	"log"
	"os"

	"github.com/jet/pkg/subcommands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:        "init",
				Usage:       "./jet init",
				Description: "Initialize a new Jet repo",
				Action: func(ctx *cli.Context) error {
					return subcommands.Init()
				},
			},
		},
	}

	// TODO: not tested yet
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
