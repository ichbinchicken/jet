package main

import (
	"github.com/jet/pkg/boundaries"
	"log"
	"os"

	"github.com/jet/pkg/subcommands"
	"github.com/urfave/cli/v2"
)

// some tools:
// https://golangci-lint.run/usage/quick-start/
func main() {
	// Q: why fileSys must to initialize with &boundaries.OsFileSystem{} rather than boundaries.OsFileSystem{}?
	// A: Short summary:
	//    An assignment to a variable of interface type is valid if the value being assigned implements the interface it is assigned to.
	//    It implements it if its method set is a superset of the interface.
	//    The method set of pointer types includes methods with both pointer and non-pointer receiver.
	//    The method set of non-pointer types only includes methods with non-pointer receiver.
	// todo:
	// https://stackoverflow.com/questions/40823315/x-does-not-implement-y-method-has-a-pointer-receiver
	fileSys := &boundaries.OsFileSystem{}
	Run(fileSys)
}

// This links explains why we don't need to have `fileSys *boundaries.FileSystem` <-- here FileSystem is an interface.
// https://stackoverflow.com/questions/13511203/why-cant-i-assign-a-struct-to-an-interface

func Run(fileSys boundaries.FileSystem) {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:        "init",
				Usage:       "./jet init",
				Description: "Initialize a new Jet repo",
				Action: func(ctx *cli.Context) error {
					// ctx is not used here
					return subcommands.Init(fileSys)
				},
			},
			{
				Name:        "commit",
				Usage:       "./jet commit",
				Description: "commit a jet change",
				Action: func(ctx *cli.Context) error {
					return subcommands.Commit(ctx, fileSys)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
