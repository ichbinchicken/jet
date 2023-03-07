package subcommands

import (
	"fmt"
	"github.com/jet/pkg/boundaries"
	"github.com/urfave/cli/v2"
)

/**
Only cares about the happy path and omit the error handling for now.
*/

func Init(ctx *cli.Context, fileSys boundaries.FileSystem) error {
	var dirs = []string{
		"objects",
		"refs",
	}

	//cwd, err := getCwd()
	//if err != nil {
	//    return fmt.Errorf("cannot get cwd: %w", err)
	//}
	//path := cwd + ctx.Args().Get(0) + ".jet"

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

//func getCwd() (string, error) {
//	ex, err := os.Executable()
//	if err != nil {
//		return "", err
//	}
//	exPath := filepath.Dir(ex)
//	return exPath, nil
//}
