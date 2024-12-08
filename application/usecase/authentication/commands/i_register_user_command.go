package commands

import "gin/api/requests"

type IRegisterUserCommand interface {
	Register(request *requests.RegisterRequest) (bool, error)
}
