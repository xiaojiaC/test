package test

import (
	pb "github.com/test/proto/example"
	"github.com/test/service"

	"github.com/stretchr/testify/assert"
)

func (s *GreeterSuite) TestSayHelloWithMissingNameParam() {
	ctx := s.GetContext()
	req := &pb.HelloRequest{
		Name: "",
	}

	_, err := greeterService.SayHello(ctx, req)

	if assert.NotNil(s.T(), err) {
		s.AssertErrorCode(err, service.MissingNameParam)
	}
}

func (s *GreeterSuite) TestSayHelloNormally() {
	ctx := s.GetContext()
	req := &pb.HelloRequest{
		Name: "World",
	}

	resp, err := greeterService.SayHello(ctx, req)

	if assert.NotNil(s.T(), resp) && assert.Nil(s.T(), err) {
		assert.Equal(s.T(), "Hello World", resp.Message)
	}
}
