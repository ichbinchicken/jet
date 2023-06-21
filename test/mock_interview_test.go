package test

import (
	"github.com/jet/pkg/zzm_mock_interview"
	"github.com/stretchr/testify/suite"
	"testing"
)

type MockOne struct {
	suite.Suite
}

func TestMockOne(t *testing.T) {
	suite.Run(t, new(MockOne))
}

func (one *MockOne) testOne() {
	zzm_mock_interview.GenerateBinaryNumber(3)
}
