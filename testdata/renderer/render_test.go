package main_test

import (
	"context"
	"testing"

	"github.com/anekkanti/go-codegen/testdata/input/service"
	"github.com/anekkanti/go-codegen/testdata/renderer/output"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

func TestRender(t *testing.T) {
	handler, err := output.NewCloudServiceServerHandler()
	assert.NoError(t, err)
	assert.NotNil(t, handler)

	_, err = handler.GetUsers(context.TODO(), &service.GetUsersRequest{})
	assert.Error(t, err)
	code, ok := status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, codes.Unimplemented, code.Code())

	_, err = handler.CreateUser(context.TODO(), &service.CreateUserRequest{})
	assert.Error(t, err)
	code, ok = status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, codes.Unimplemented, code.Code())

}
