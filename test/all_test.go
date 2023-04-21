package test

import (
	"github.com/jet/pkg"
	"github.com/jet/pkg/boundaries"
	"github.com/jet/pkg/helper"
	"github.com/jet/pkg/subcommands"
	"github.com/stretchr/testify/suite"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
	"testing"
)

type SequentialOperationsTestSuite struct {
	fileSys       *boundaries.OsFileSystem
	targetDirPath string
	app           *cli.App
	suite.Suite
}

func (suite *SequentialOperationsTestSuite) SetupTest() {
	var err error
	suite.fileSys = &boundaries.OsFileSystem{}
	suite.app = app.NewCliApp(suite.fileSys)
	suite.targetDirPath, err = helper.AnyToString(suite.app.Metadata["targetDirPath"])
	if err != nil {
		panic(err)
	}

	if strings.HasSuffix(suite.T().Name(), "_SkipSetup") {
		return
	}

	err = os.RemoveAll(suite.targetDirPath + "./jet")
	if err != nil {
		panic(err)
	}

	err = suite.app.Run([]string{"program", "init"})
	if err != nil {
		panic(err)
	}
}

func (suite *SequentialOperationsTestSuite) TestCommitOnce() {
	err := suite.app.Run([]string{"program", "commit"})
	if err != nil {
		panic(err)
	}
}

// commit_action.go
func (suite *SequentialOperationsTestSuite) TestGetDirEntriesWithoutJet_SkipSetup() {
	entries := subcommands.GetDirEntriesWithoutJet(suite.targetDirPath, suite.fileSys)
	suite.NotContains(entries, ".jet")
}

func TestSequentialOperations(t *testing.T) {
	suite.Run(t, new(SequentialOperationsTestSuite))
}
