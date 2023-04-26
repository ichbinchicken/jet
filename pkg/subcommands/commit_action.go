package subcommands

import (
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
	var blobs []object.Blob
	// TODO: right now we only care about files only
	for _, entry := range GetDirEntriesWithoutJet(rootPath, fs) {
		// write blob objects
		if !entry.IsDir() {
			contents, err := fs.Read(filepath.Join(rootPath, entry.Name()))
			if err != nil {
				return err
			}
			blob := object.NewBlob(contents, entry.Name())
			blobs = append(blobs, blob)
			err = fs.WriteJetObject(objectsPath, &blob)
			if err != nil {
				return err
			}
		}
	}
	// write tree objects based on blobs
	//tree := object.NewTree(blobs)
	//err := fs.WriteJetObject(objectsPath, &tree)
	//if err != nil {
	//	return err
	//}

	return nil
}

func GetDirEntriesWithoutJet(path string, fs database.FileStorage) []os.DirEntry {
	var dirEntries helper.GenericSlice[os.DirEntry]
	var err error
	dirEntries, err = fs.ListFiles(path)
	if err != nil {
		panic(err)
	}

	dirEntries = dirEntries.FilterOut(func(e os.DirEntry) bool {
		if e.Name() == helper.DOTJET {
			return true
		} else {
			return false
		}
	})

	return dirEntries
}
