package contracts

import (
	"gin/api/requests"
	"gin/application/utility"
)

type ILogOutCommand interface {
	// LogOut is a method to log out user
	LogOut(request *requests.LogOutRequest) (bool, *utility.ErrorCode)
}
