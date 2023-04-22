package app

import (
	"github.com/jet/pkg/database"
	"github.com/jet/pkg/subcommands"
	"github.com/urfave/cli/v2"
)

/*
This links explains why we don't need to have `fs *database.FileStorage` <-- here FileStorage is an interface.
https://stackoverflow.com/questions/13511203/why-cant-i-assign-a-struct-to-an-interface
*/
func NewCliApp(fs database.FileStorage) *cli.App {
	return &cli.App{
		Metadata: map[string]interface{}{
			"targetDirPath": "/Users/zzheng2/glacier/jet-sample-repo",
		},
		Commands: []*cli.Command{
			{
				Name:        "init",
				Usage:       "./jet init",
				Description: "Initialize a new Jet repo",
				Action: func(ctx *cli.Context) error {
					return subcommands.Init(ctx, fs)
				},
			},
			{
				Name:        "commit",
				Usage:       "./jet commit",
				Description: "commit a jet change",
				Action: func(ctx *cli.Context) error {
					return subcommands.Commit(ctx, fs)
				},
			},
		},
	}
}
