package subcommands

import (
	"fmt"
	"github.com/jet/pkg/database"
	"github.com/jet/pkg/helper"
	"github.com/urfave/cli/v2"
	"path/filepath"
)

/**
 * Only cares about the happy path and omit the error handling for now.
 */

func Init(ctx *cli.Context, fs database.FileStorage) error {
	var dirs [2]string
	dirs[0] = helper.OBJECTS
	dirs[1] = helper.REFS

	rootPath := helper.AnyToString(ctx.App.Metadata["targetDirPath"])
	for _, dir := range dirs {
		err := fs.MakeDirs(filepath.Join(rootPath, helper.DOTJET, dir))
		if err != nil {
			// todo: understand what does %w do???
			return fmt.Errorf("cannot create a new dir: %w", err)
		}
	}

	fmt.Println("Initialized empty jet repo")

	return nil
}
