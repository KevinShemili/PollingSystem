package contracts

import "gin/api/requests"

type IRegisterCommand interface {
	Register(request *requests.RegisterRequest) (bool, error)
}
