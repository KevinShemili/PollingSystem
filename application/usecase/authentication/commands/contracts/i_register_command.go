package contracts

import (
	"gin/api/requests"
	"gin/application/utility"
)

type IRegisterCommand interface {
	Register(request *requests.RegisterRequest) (bool, *utility.ErrorCode)
}
