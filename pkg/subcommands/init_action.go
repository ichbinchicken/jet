package subcommands

import (
	"fmt"
	"github.com/jet/pkg/boundaries"
)

/**
 * Only cares about the happy path and omit the error handling for now.
 */

func Init(fileSys boundaries.FileSystem) error {
	var dirs = []string{
		"objects",
		"refs",
	}

	// It will just create a .jet folder at the CWD.
	for _, dir := range dirs {
		err := fileSys.MakeDirs("./.jet/" + dir)
		if err != nil {
			// todo: understand what does %w do???
			return fmt.Errorf("cannot create a new dir: %w", err)
		}
	}

	fmt.Printf("Initialized empty jet repo")

	return nil
}
