package user

type Access struct {
	// The account access
	AccountAccess string `protobuf:"bytes,1,opt,name=account_access,json=accountAccess,proto3" json:"account_access,omitempty"`
}

type UserSpec struct {
	// The email address associated to the user
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	// The access to assigned to the user
	Access *Access `protobuf:"bytes,2,opt,name=access,proto3" json:"access,omitempty"`
}

type User struct {
	// The id of the user
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The user specification
	Spec *UserSpec `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec,omitempty"`
	// The current state of the user
	State string `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
}
