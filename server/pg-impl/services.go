package pgimpl

import "context"

type ProductService interface {

	// GetLimitaryHour needs the following required input:
	//  - ProductType
	// The returned USER_ERROR would be as below:
	//  - Code_INSUFFICIENT_REQUEST
	GetLimitaryHour(context.Context, GetLimitaryHourRequest) (GetLimitaryHourReply, error)
}
