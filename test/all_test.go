package test

import (
	"fmt"
	"github.com/jet/pkg"
	"github.com/jet/pkg/database"
	"github.com/jet/pkg/helper"
	"github.com/jet/pkg/subcommands"
	"github.com/jet/pkg/subcommands/object"
	"github.com/stretchr/testify/suite"
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type AllTest struct {
	fs            *database.OsFileStorage
	targetDirPath string
	app           *cli.App
	suite.Suite
}

func (suite *AllTest) SetupTest() {
	var err error
	suite.fs = &database.OsFileStorage{}
	suite.app = app.NewCliApp(suite.fs)
	suite.targetDirPath = helper.AnyToString(suite.app.Metadata["targetDirPath"])

	if strings.HasSuffix(suite.T().Name(), "_SkipSetup") {
		return
	}

	err = os.RemoveAll(filepath.Join(suite.targetDirPath, helper.DOTJET))
	if err != nil {
		panic(err)
	}

	err = suite.app.Run([]string{"program", "init"})
	if err != nil {
		panic(err)
	}
}

func (suite *AllTest) TestCommitOnce() {
	err := suite.app.Run([]string{"program", "commit"})

	if err != nil {
		panic(err)
	}
}

// commit_action.go
func (suite *AllTest) TestGetDirEntriesWithoutJet_SkipSetup() {
	entries := subcommands.GetDirEntriesWithoutJet(suite.targetDirPath, suite.fs)
	suite.NotContains(entries, helper.DOTJET)
}

// object.go
func (suite *AllTest) TestSHA1_SkipSetup() {
	s := object.GenerateSHA1Hash("abc")
	fmt.Println(string(s))
}

// test writeBlob
func (suite *AllTest) TestWriteBlob_SkipSetup() {
	blob := object.NewBlob("hello world 你好世界")
	err := suite.fs.WriteBlob(filepath.Join(suite.targetDirPath), &blob)
	defer os.RemoveAll(filepath.Join(suite.targetDirPath, blob.ObjId()[:2]))
	suite.NoError(err)
	compressed, err := suite.fs.Read(filepath.Join(suite.targetDirPath, blob.ObjId()[:2], blob.ObjId()[2:]))
	suite.NoError(err)
	out, err := helper.Decompress([]byte(compressed))
	suite.NoError(err)
	suite.Equal(fmt.Sprintf("blob 24\x00hello world 你好世界"), string(out))
}

func TestSequentialOperations(t *testing.T) {
	suite.Run(t, new(AllTest))
}
