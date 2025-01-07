package contracts

import (
	"gin/api/requests"
	"gin/application/utility"
)

type IRegisterCommand interface {
	// Register is a method to register a user
	Register(request *requests.RegisterRequest) (bool, *utility.ErrorCode)
}
