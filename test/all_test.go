package test

import (
	"encoding/hex"
	"fmt"
	"github.com/jet/pkg"
	"github.com/jet/pkg/database"
	"github.com/jet/pkg/helper"
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

// blob.go
func (suite *AllTest) TestSHA1_SkipSetup() {
	out := object.GenerateSHA1Hash("abc")
	fmt.Println(hex.EncodeToString(out))
}

func (suite *AllTest) TestWriteBlobUnit_SkipSetup() {
	original, err := suite.fs.Read(filepath.Join(".", "testing_file_ascii_only.txt"))
	suite.NoError(err)

	b := object.NewBlob(original, "testing_file_ascii_only.txt")
	hexStr := hex.EncodeToString(b.Oid())

	err = suite.fs.WriteJetObject(".", &b)
	suite.NoError(err)

	compressed, err := os.ReadFile(filepath.Join(".", hexStr[:2], hexStr[2:]))
	raw, err := helper.Decompress(compressed)
	suite.NoError(err)
	suite.Equal(fmt.Sprintf("blob 2003\x00%s", original), string(raw))

	// clean up the created blob
	err = os.RemoveAll(filepath.Join(".", hexStr[:2]))
	suite.NoError(err)
}

func (suite *AllTest) TestWriteTreeUnit_SkipSetup() {
	blob := object.NewBlob("hello world 你好世界", "dummy.txt")
	blob2 := object.NewBlob("こんにちは aloha", "dummy2.rb")
	blob3 := object.NewBlob("こんにちは aloha", "dummy3.go")
	blobs := []object.Blob{blob, blob2, blob3}
	tree := object.NewTree(blobs)
	hexStr := hex.EncodeToString(tree.Oid())

	err := suite.fs.WriteJetObject(".", &tree)
	suite.NoError(err)

	compressed, err := os.ReadFile(filepath.Join(".", hexStr[:2], hexStr[2:]))
	suite.NoError(err)
	raw, err := helper.Decompress(compressed)
	suite.NoError(err)
	suite.Equal(tree.Odata(), raw)

	// clean up the created blob
	err = os.RemoveAll(filepath.Join(".", hexStr[:2]))
	suite.NoError(err)
}

func TestSequentialOperations(t *testing.T) {
	suite.Run(t, new(AllTest))
}
