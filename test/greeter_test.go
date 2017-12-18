package test

import (
	"testing"

	pb "github.com/test/proto/example"
	"github.com/test/service"

	"github.com/stretchr/testify/suite"
)

var greeterService pb.GreeterServer

type GreeterSuite struct {
	BaseSuite
}

func init() {
	greeterService = new(service.GreeterService)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run .
func TestGreeterSuite(t *testing.T) {
	greeterService := new(GreeterSuite)
	suite.Run(t, greeterService)
}

func (s *GreeterSuite) SetupSuite() {
	s.Setup()

	// Default mock data construction
}

func (s *GreeterSuite) TearDownSuite() {
	s.TearDown()
}
