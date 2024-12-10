package contracts

import (
	"gin/api/requests"
	"gin/application/usecase/authentication/results"
	"gin/application/utility"
)

type IRefreshCommand interface {
	Refresh(request *requests.TokensRequest) (*results.RefreshResult, *utility.ErrorCode)
}
