package test

import (
	"github.com/test/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"golang.org/x/net/context"
)

type BaseSuite struct {
	suite.Suite
}

// Setup inits the suite test environment.
func (suite *BaseSuite) Setup() {
}

// TearDown clean up the suite test environment.
func (suite *BaseSuite) TearDown() {

}

func (suite *BaseSuite) GetContext() context.Context {
	return context.Background()
}

func (suite *BaseSuite) AssertErrorCode(err error, code service.Code) {
	if assert.NotNil(suite.T(), err) {
		if rpcError, ok := err.(*service.RPCError); ok {
			if assert.NotNil(suite.T(), rpcError) {
				assert.Equal(suite.T(), code, rpcError.Code)
			}
		}
	}
}
