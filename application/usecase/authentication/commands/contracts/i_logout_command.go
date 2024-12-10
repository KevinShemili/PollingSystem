package contracts

import (
	"gin/api/requests"
	"gin/application/utility"
)

type ILogOutCommand interface {
	LogOut(request *requests.LogOutRequest) (bool, *utility.ErrorCode)
}
