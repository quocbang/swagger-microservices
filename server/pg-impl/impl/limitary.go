package impl

import (
	"context"
	"errors"
	pgimpl "quocbang/swagger-microservices/pg-impl"
	"quocbang/swagger-microservices/pg-impl/impl/orm/models"
	"time"

	"gorm.io/gorm"
)

// CreateBatch implements gitlab.kenda.com.tw/kenda/mcom DataManager interface.
func (dm *DataManager) GetLimitaryHour(ctx context.Context, req pgimpl.GetLimitaryHourRequest) (pgimpl.GetLimitaryHourReply, error) {

	var limiatry models.LimitaryHour
	session := dm.newSession(ctx)

	if err := session.db.Where(&models.LimitaryHour{ProductType: req.ProductType}).Take(&limiatry).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return pgimpl.GetLimitaryHourReply{LimitaryHour: pgimpl.LimitaryHourReply{
				Min: time.Now().Add(0),
				Max: time.Now().Add(0),
			}}, nil
		}
	}
	return pgimpl.GetLimitaryHourReply{LimitaryHour: pgimpl.LimitaryHourReply{
		Min: time.Now().Add(time.Duration(limiatry.Min) * time.Hour),
		Max: time.Now().Add(time.Duration(limiatry.Max) * time.Hour),
	}}, nil
}
