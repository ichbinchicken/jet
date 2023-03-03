package subcommands

import (
    "github.com/stretchr/testify/suite"
    "testing"
)

// TODO: learn how to use golang testify mock to mock file system and log
type JetCommandsTestSuite struct {
    suite.Suite
}
func (jetSuite *JetCommandsTestSuite) SetupTest() {

}

func TestJetCommandsSuite(t *testing.T) {
    // this will initialize JetCommandsTestSuite.Suite instance
    suite.Run(t, new(JetCommandsTestSuite))
}