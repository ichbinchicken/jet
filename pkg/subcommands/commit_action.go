package subcommands

import (
	"fmt"
	"github.com/jet/pkg/boundaries"
	"github.com/urfave/cli/v2"
)

func Commit(ctx *cli.Context, fileSys boundaries.FileSystem) error {
	path := ctx.App.Metadata["targetDirPath"].(string)
	fmt.Println(path)
	dirEntries, err := fileSys.ListFiles(".")
	if err != nil {
		return err
	}

	fmt.Println(len(dirEntries))
	for _, e := range dirEntries {
		fmt.Println(e.Name())
	}

	return nil
}
