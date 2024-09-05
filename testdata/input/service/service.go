package service

import "context"

type CloudServiceServer interface {
	// Gets all known users
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error)
	// Create a user
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
}
