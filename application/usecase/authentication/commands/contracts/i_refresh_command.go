package contracts

import (
	"gin/api/requests"
	"gin/application/usecase/authentication/results"
	"gin/application/utility"
)

type IRefreshCommand interface {
	// Refresh is a method to generate new JWT & refresh token
	Refresh(request *requests.TokensRequest) (*results.RefreshResult, *utility.ErrorCode)
}
