package pgimpl

import (
	"fmt"
	"time"
)

type LimitaryHour struct {
	ProductType  string
	LimitaryHour LimitaryHourParameter
}

type LimitaryHourParameter struct {
	Min int32
	Max int32
}

type CreateLimitaryHourRequest struct {
	LimitaryHour []LimitaryHour
}

func (req CreateLimitaryHourRequest) CheckInsufficiency() error {
	if len(req.LimitaryHour) == 0 {
		return fmt.Errorf(fmt.Sprintf("product type can't be empty %v", req.LimitaryHour))
	}
	// ProductType cant be empty, Min or Max never negative number,Max cannot be less than Min.
	for _, v := range req.LimitaryHour {
		if v.ProductType == "" ||
			v.LimitaryHour.Min < 0 ||
			v.LimitaryHour.Max < 0 ||
			v.LimitaryHour.Max < v.LimitaryHour.Min {
			return fmt.Errorf(fmt.Sprintf("insufficiency %v", req.LimitaryHour))
		}
	}
	return nil
}

type GetLimitaryHourRequest struct {
	ProductType string
}

type LimitaryHourReply struct {
	Min time.Time
	Max time.Time
}

type GetLimitaryHourReply struct {
	LimitaryHour LimitaryHourReply
}

// CheckInsufficiency implements gitlab.kenda.com.tw/kenda/mcom Request interface.
// todo delete after rewriting Request interface
func (req GetLimitaryHourRequest) CheckInsufficiency() error {
	return nil
}
