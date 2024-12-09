package contracts

import (
	"gin/api/requests"
	"gin/application/usecase/authentication/results"
	"gin/application/utility"
)

type ILoginCommand interface {
	Login(request *requests.LoginRequest) (*results.LoginResult, *utility.ErrorCode)
}
