package commit

import (
	"encoding/hex"
	"github.com/jet/pkg/database"
	"github.com/jet/pkg/helper"
	"github.com/jet/pkg/subcommands/object"
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
	"time"
)

func Commit(ctx *cli.Context, fs database.FileStorage) error {
	rootPath := helper.AnyToString(ctx.App.Metadata["targetDirPath"])
	objectsPath := filepath.Join(rootPath, helper.DOTJET, helper.OBJECTS)
	var blobs []object.Blob

	// TODO: right ti we only care about files only
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

	// write tree object based on blobs:
	tree := object.NewTree(blobs)
	err := fs.WriteJetObject(objectsPath, &tree)
	if err != nil {
		return err
	}

	// write commit object based on tree object:
	// TODO: read GIT_AUTHOR_NAME & GIT_AUTHOR_EMAIL env var
	author := object.NewAuthor("Zheng Ziming", "zzm@jet.com", time.Now())
	commitMsg := helper.ReadStdin()
	commit := object.NewCommit(author, commitMsg, tree.Oid())
	err = fs.WriteJetObject(objectsPath, &commit)
	if err != nil {
		return err
	}

	// update .jet/HEAD
	updateHEAD(filepath.Join(objectsPath, "..", "HEAD"), hex.EncodeToString(commit.Oid()))

	return nil
}

func updateHEAD(path string, commitHexStr string) {
	_, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(path, []byte(commitHexStr), 0o644)
	if err != nil {
		panic(err)
	}
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
