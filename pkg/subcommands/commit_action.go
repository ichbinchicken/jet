package subcommands

import (
	"github.com/jet/pkg/boundaries"
	"github.com/jet/pkg/helper"
	"github.com/urfave/cli/v2"
	"os"
)

func Commit(ctx *cli.Context, fileSys boundaries.FileSystem) error {
	_, err := helper.AnyToString(ctx.App.Metadata["targetDirPath"])
	if err != nil {
		panic(err)
	}

	//GetDirEntriesWithoutJet(path, fileSys)

	return nil
}

func GetDirEntriesWithoutJet(path string, fileSys boundaries.FileSystem) []os.DirEntry {
	var dirEntries helper.GenericSlice[os.DirEntry]
	var err error
	dirEntries, err = fileSys.ListFiles(path)
	if err != nil {
		panic(err)
	}

	/*
			@zzmlearning
			try to use filter()-like function so that I could do
		    return dirEntries.filter(entry -> entry.Name() != ".jet")
	*/

	dirEntries = dirEntries.FilterOut(func(e os.DirEntry) bool {
		if e.Name() == ".jet" {
			return true
		} else {
			return false
		}
	})

	return dirEntries
}
