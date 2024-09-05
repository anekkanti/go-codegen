package service

import "github.com/anekkanti/go-codegen/testdata/input/user"

type GetUsersRequest struct {
	// The requested size of the page to retrieve - optional.
	// Cannot exceed 1000. Defaults to 100.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The page token if this is continuing from another response - optional.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Filter users by email address - optional.
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	// Filter users by the namespace they have access to - optional.
	Namespace string `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

type GetUsersResponse struct {
	// The list of users in ascending ids order
	Users []*user.User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	// The next page's token
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

type CreateUserRequest struct {
	// The spec for the user to invite
	Spec *user.UserSpec `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	// The id to use for this async operation - optional
	AsyncOperationId string `protobuf:"bytes,2,opt,name=async_operation_id,json=asyncOperationId,proto3" json:"async_operation_id,omitempty"`
}

type CreateUserResponse struct {
	// The id of the user that was invited
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// The async operation
	AsyncOperationId *string `protobuf:"bytes,2,opt,name=async_operation_id,json=asyncOperation,proto3" json:"async_operation_id,omitempty"`
}
