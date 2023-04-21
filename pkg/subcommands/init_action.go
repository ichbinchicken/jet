package subcommands

import (
	"fmt"
	"github.com/jet/pkg/boundaries"
	"github.com/jet/pkg/helper"
	"github.com/urfave/cli/v2"
)

/**
 * Only cares about the happy path and omit the error handling for now.
 */

func Init(ctx *cli.Context, fileSys boundaries.FileSystem) error {
	var dirs = []string{
		"objects",
		"refs",
	}

	path, err := helper.AnyToString(ctx.App.Metadata["targetDirPath"])
	if err != nil {
		panic(err)
	}
	for _, dir := range dirs {
		err := fileSys.MakeDirs(path + "/.jet/" + dir)
		if err != nil {
			// todo: understand what does %w do???
			return fmt.Errorf("cannot create a new dir: %w", err)
		}
	}

	fmt.Println("Initialized empty jet repo")

	return nil
}
