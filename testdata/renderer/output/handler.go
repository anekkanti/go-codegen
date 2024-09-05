// Code generated by go-codegen. DO NOT EDIT.
package output

import (
	"context"
	service "github.com/anekkanti/go-codegen/testdata/input/service"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// handler for CloudServiceServer
type CloudServiceServerHandler struct{}

// create a new CloudServiceServerHandler
func NewCloudServiceServerHandler() (service.CloudServiceServer, error) {
	return &CloudServiceServerHandler{}, nil
}

func (h *CloudServiceServerHandler) GetUsers(ctx context.Context, req *service.GetUsersRequest) (*service.GetUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented")
}

func (h *CloudServiceServerHandler) CreateUser(ctx context.Context, req *service.CreateUserRequest) (*service.CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented")
}
