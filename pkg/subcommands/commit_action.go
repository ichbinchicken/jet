package subcommands

import (
	"fmt"
	"github.com/jet/pkg/database"
	"github.com/jet/pkg/helper"
	"github.com/jet/pkg/subcommands/object"
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
)

func Commit(ctx *cli.Context, fs database.FileStorage) error {
	rootPath := helper.AnyToString(ctx.App.Metadata["targetDirPath"])
	objectsPath := filepath.Join(rootPath, helper.DOTJET, helper.OBJECTS)

	// creating blobs
	for _, entry := range GetDirEntriesWithoutJet(rootPath, fs) {
		// right now we only create a blob when it's a file
		if !entry.IsDir() {
			contents, err := fs.Read(filepath.Join(rootPath, entry.Name()))
			if err != nil {
				return err
			}
			blob := object.NewBlob(contents)
			fmt.Println(blob.ObjId())
			fs.WriteBlob(objectsPath, &blob)
		}
	}

	return nil
}

//func createBlob(fs database.FileStorage, path string) *object.Object {
//}

func GetDirEntriesWithoutJet(path string, fs database.FileStorage) []os.DirEntry {
	var dirEntries helper.GenericSlice[os.DirEntry]
	var err error
	dirEntries, err = fs.ListFiles(path)
	if err != nil {
		panic(err)
	}

	/*
			@zzmlearning
			try to use filter()-like function so that I could do
		    return dirEntries.filter(entry -> entry.Name() != ".jet")
	*/

	dirEntries = dirEntries.FilterOut(func(e os.DirEntry) bool {
		if e.Name() == helper.DOTJET {
			return true
		} else {
			return false
		}
	})

	return dirEntries
}
